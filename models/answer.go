package models

type Answer struct {
	UserID    int    `gorm:"column:user_id"`
	Email     string `json:"email"`
	Question1 int    `gorm:"column:questionid1"`
	Answer1   string `json:"answer1"`
	Question2 int    `gorm:"column:questionid2"`
	Answer2   string `json:"answer2"`
	Question3 int    `gorm:"column:questionid3"`
	Answer3   string `json:"answer3"`
}

type Question struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Content string `json:"content"`
}
