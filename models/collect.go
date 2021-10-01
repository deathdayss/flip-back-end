package models

type Collect struct {
	ID     int `gorm:"column:id;AUTO_INCREMENT"`
	UserID int `gorm:"column:user_id"`
	GameID int `gorm:"column:game_id"`
}
