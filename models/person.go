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
}

type UserClaims struct {
	Email string
	jwt.StandardClaims
}

type PersonImg struct {
	ID  int    `gorm:"primary_key;AUTO_INCREMENT"`
	UID int    `gorm:"uid; not null"`
	URL string `gorm:"url; not null`
}

type PersonDetail struct {
	Email  string `gorm:"email; not null; unique"`
	Birth  string `gorm:"birth"`
	Age    int    `gorm:"age"`
	Gender string `gorm: "gender"`
}

type Author struct {
	URL      string `gorm:"url; not null`
	NickName string `json:"nickname"`
	LikeNum  int    `json:"like_num"`
}
