package models

type Answer struct {
	UserID    int    `gorm:"column:user_id"`
	Email     string `json:"email"`
	Question1 int    `gorm:"column:question1id"`
	Answer1   string `json:"answer1"`
	Question2 int    `gorm:"column:question2id"`
	Answer2   string `json:"answer2"`
	Question3 int    `gorm:"column:question3id"`
	Answer3   string `json:"answer3"`
}

type Question struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Content string `json:"content"`
}
