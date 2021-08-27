package models

type Person struct {
    ID int `gorm:"primary_key;AUTO_INCREMENT"`
    Email string `gorm:"unique;not null"`
    Nickname string `json:"nickname"`
    Password string `json:"password"`
}
