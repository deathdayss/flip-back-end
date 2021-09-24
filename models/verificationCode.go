package models

type Code struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Content string `json:"content"`
	URL     string `json:"url"`
}
