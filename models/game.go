package models

import "time"

type Game struct {
	ID           int       `gorm:"column:id;AUTO_INCREMENT"`
	Name         string    `json:"name"`
	Introduction string    `json:"introduction"`
	ImgUrl       string    `json:"img_url"`
	UID          int       `json:"uid"`
	FileUrl      string    `json:"file_url"`
	Zone         string    `json:"zone"`
	CreateAt     time.Time `json:"create_at"`
}

/*
调整Game表结构之前的版本
type Game struct {
	ID          int    `gorm:"column:id;AUTO_INCREMENT"`
	Name        string `json:"name"`
	LikeNum     int    `json:"like_num"`
	DownloadNum int    `json:"download_num"`
	CommentNum  int    `json:"comment_num"`
	ImgUrl      string `json:"img_url"`
	UID         int    `json:"uid"`
	FileUrl     string `json:"file_url"`
	Zone        string `json:"zone"`
}

*/
