package models

type Answer struct {
	UserID   int    `gorm:"column:user_id"`
	Email    string `json:"email"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
