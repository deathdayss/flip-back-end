package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Person struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Email    string `gorm:"unique;not null"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
	Person
}

type UserClaims struct {
	Email string
	jwt.StandardClaims
}
