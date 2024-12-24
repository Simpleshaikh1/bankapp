package migrations

import (
	"database/sql"
	"fintechApp/helpers"
	"fintechApp/interfaces"
)

func createAccount() {
	db := helpers.ConnectDB()

	postgreSQLDB, err := db.DB()
	helpers.HandleError(err, "Failed to connect to database")

	defer func(postgreSQLDB *sql.DB) {
		err := postgreSQLDB.Close()
		helpers.HandleError(err, "Failed to close database connection")
	}(postgreSQLDB)

	users := &[2]interfaces.User{
		{UserName: "Toyyib", Email: "simpleshaikh@gmail.com"},
		{UserName: "Ahmad", Email: "ahmad@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		usersEmail := []byte(users[i].Email)
		generatePassword := helpers.HashAndSalt(usersEmail)
		user := &interfaces.User{UserName: users[i].UserName, Email: users[i].Email, Password: generatePassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].UserName + "s" + " account"),
			Balance: uint(1000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	postgreSQLDB, err := db.DB()
	helpers.HandleError(err, "Failed to connect to database")

	defer func(postgreSQLDB *sql.DB) {
		err := postgreSQLDB.Close()
		helpers.HandleError(err, "Failed to close database connection")
	}(postgreSQLDB)

	err = db.AutoMigrate(User, Account)
	if err != nil {
		return
	}

	createAccount()
}
