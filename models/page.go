package models

type Page struct {
	Id        int    `gorm:"primary_key;not null"`
	Name      string `json:"name"`
	Imagepath string `json:"imagepath"`
}
