package dto

type ProductItem struct {
	Name    string `json:"game_name"`
	LikeNum int    `json:"like_num"`
	//CollectionNum int `json:"collection_num"`
	//ShareNum 		int `json:"share_num"`
	//Introduction  String `json:"introduction_num"`
	Img 	[]byte `json:"img"`
	UID		int 	`json:uid`
	//Time
	//ClickCount
}
