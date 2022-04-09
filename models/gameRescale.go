package models

type GameRescale struct {
	ID          int    `gorm:"column:id;AUTO_INCREMENT"`
	Name        string `json:"name"`
	LikeNum     int    `json:"like_num"`
	DownloadNum int    `json:"download_num"`
	CommentNum  int    `json:"comment_num"`
	ImgUrl      string `json:"img_url"`
	UID         int    `json:"uid"`
	FileUrl     string `json:"file_url"`
}
