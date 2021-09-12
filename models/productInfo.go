package models

type ProductInfo struct {
	ID            int    `gorm:"column:id;AUTO_INCREMENT"`
	Name          string `json:"game_name"`
	LikeNum       int    `json:"like_num"`
	CollectionNum int    `json:"collection_num"`
	ShareNum      int    `json:"share_num"`
	Introduction  string `json:"introduction"`
	ImgUrl        string `json:"img_url"`
	UID           int    `json:"uid"`
	Time          string `json:"time"`
	ClickCount    int    `json:"click_count"`
}
