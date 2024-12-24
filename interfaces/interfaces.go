package interfaces

import "gorm.io/gorm"

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

type ResponseAccount struct {
	ID      uint
	Balance int
	Name    string
}

type ResponseUser struct {
	ID       uint
	UserName string
	Email    string
	Accounts []ResponseAccount
}
