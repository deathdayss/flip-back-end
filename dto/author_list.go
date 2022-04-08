package dto

type AuthorItem struct {
	URL      string `gorm:"url; not null`
	NickName string `json:"nickname"`
	LikeNum  int    `json:"like_num"`
}
