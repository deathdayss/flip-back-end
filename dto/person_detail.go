package dto

type PersonDetail struct {
	Email  string `json:email example:"123@123.com"` 
	Age    int    `json:age example:"12"`
	Gender string `json:gender example:"female"`
	Birth  string    `json:birth example:"1992-12-10"`
}