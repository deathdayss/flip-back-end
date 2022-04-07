package dto

type RankItem struct {
	ID          int    `json:"GID"`
	Name        string `json:"game_name"`
	LikeNum     int    `json:"like_num"`
	DownloadNum int    `json:"download_num`
	CommentNum  int    `json:"comment_num`
	Img         string `json:"img"`
	AuthorName  string `json:author_name`
}

type RankItemByZone struct {
	ID          int    `json:"GID"`
	Name        string `json:"game_name"`
	LikeNum     int    `json:"like_num"`
	DownloadNum int    `json:"download_num`
	CommentNum  int    `json:"comment_num`
	Img         string `json:"img"`
	AuthorName  string `json:author_name`
	ClickCount  int    `json:click_count`
}
