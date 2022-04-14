package models

type Answer struct {
	UserID    int    `gorm:"column:user_id"`
	Email     string `json:"email"`
	Question1 int    `gorm:"question1id; not null"`
	Answer1   string `json:"answer1"`
	Question2 int    `gorm:"question1id; not null"`
	Answer2   string `json:"answer2"`
	Question3 int    `gorm:"question1id; not null"`
	Answer3   string `json:"answer3"`
}

type Question struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Content string `json:"content"`
}
