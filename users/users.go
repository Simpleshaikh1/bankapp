package users

import (
	"fintechApp/helpers"
	"fintechApp/interfaces"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Login(username string, password string) map[string]interface{} {
	//Connect to db
	db := helpers.ConnectDB()
	user := &interfaces.User{}
	if db.Where("user_name = ?", username).First(&user).Error != nil {
		return map[string]interface{}{"message": "User not found"}
	}
	//verify password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Password is incorrect"}
	}

	//find account for user
	accounts := []interfaces.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

	//setup response
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Accounts: accounts,
	}
	//sign token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleError(err, "Failed to sign token")

	//Prepare Response
	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser
	return response
}
