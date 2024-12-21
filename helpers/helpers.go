package helpers

import (
	"golang.org/x/crypto/bcrypt"
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
