package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func HandleError(err error, message string) {
	if err != nil {
		panic(message + ": " + err.Error())
	}
}

func HashAndSalt(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	HandleError(err, "Failed to hash password")

	return string(hashed)
}

func ConnectDB() *gorm.DB {
	dsn := "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	HandleError(err, "Failed to connect to database")

	return db
}
