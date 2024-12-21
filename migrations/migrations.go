package migrations

import (
	"database/sql"
	"fintechApp/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	dsn := "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.HandleError(err, "Failed to connect to database")

	return db
}

func createAccount() {
	db := connectDB()

	postgreSQLDB, err := db.DB()
	helpers.HandleError(err, "Failed to connect to database")

	defer func(postgreSQLDB *sql.DB) {
		err := postgreSQLDB.Close()
		helpers.HandleError(err, "Failed to close database connection")
	}(postgreSQLDB)

	users := [2]User{
		{UserName: "Toyyib", Email: "simpleshaikh@gmail.com"},
		{UserName: "Ahmad", Email: "ahmad@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		usersEmail := []byte(users[i].Email)
		generatePassword := helpers.HashAndSalt(usersEmail)
		user := User{UserName: users[i].UserName, Email: users[i].Email, Password: generatePassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].UserName + "s" + " account"),
			Balance: uint(1000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	db := connectDB()
	postgreSQLDB, err := db.DB()
	helpers.HandleError(err, "Failed to connect to database")

	defer func(postgreSQLDB *sql.DB) {
		err := postgreSQLDB.Close()
		helpers.HandleError(err, "Failed to close database connection")
	}(postgreSQLDB)

	err = db.AutoMigrate(&User{}, &Account{})
	if err != nil {
		return
	}

	createAccount()
}
