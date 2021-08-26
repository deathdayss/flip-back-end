package models

type Person struct {
    Email string `gorm:"primary_key;not null"`
    Nickname string `json:"nickname"`
    Password string `json:"password"`
}
