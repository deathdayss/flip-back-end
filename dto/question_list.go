package dto

type QuestionItem struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Content string `json:"content"`
}
