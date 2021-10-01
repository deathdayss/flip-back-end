package models

type Share struct {
	ID       int `gorm:"column:id;AUTO_INCREMENT"` //这个并不是GameID？
	GameID   int `gorm:"column:game_id"`
	ShareNum int `gorm:"column:share_num"`
}
