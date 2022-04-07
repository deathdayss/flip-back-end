package models

type Click struct {
	ID     int `gorm:"column:id;AUTO_INCREMENT"`
	GameID int `gorm:"column:game_id"`
	Count  int `gorm:"column:count"`
}
