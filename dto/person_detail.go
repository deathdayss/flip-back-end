package dto

type PersonDetail struct {
	Email  string `json:email`
	Age    int    `json:age`
	Gender string `json:gender`
	Birth  string    `json:birth`
}