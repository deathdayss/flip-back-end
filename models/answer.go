package models

type Answer struct {
	UserID    int    `gorm:"column:user_id"`
	Email     string `json:"email"`
	Question1 string `json:"question1"`
	Answer1   string `json:"answer1"`
	Question2 string `json:"question2"`
	Answer2   string `json:"answer2"`
	Question3 string `json:"question3"`
	Answer3   string `json:"answer3"`
}
