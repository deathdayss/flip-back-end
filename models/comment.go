package models

import "time"

type Comment struct {
	ID         int `gorm:"column:id;AUTO_INCREMENT"`
	GID        int `gorm:"not null"`
	UID        int `gorm:"not null"`
	CreateTime time.Time
	Content string `gorm:"not null"`
	Up int
	Down int
}

type CommentUp struct {
	CID        int `gorm:"not null;primaryKey"`
	UID int `gorm:"not null;primaryKey"`
}