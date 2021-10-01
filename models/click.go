package models

import "time"

type Click struct {
	ID         int       `gorm:"column:id;AUTO_INCREMENT"`
	UserID     int       `gorm:"column:user_id"`
	GameID     int       `gorm:"column:game_id"`
	Date       time.Time `gorm:"column:date"`
	ClickCount int       `gorm:"column:click_count"`
}
