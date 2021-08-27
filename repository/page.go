package repository

import (
	"fmt"
)

//insert data
func Initpagedata() {
	pages := []models.Page{}

	session := DbEngine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, page := range pages {
		_, error = session.Insert(&page)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
}

//rank data

func Rank() []models.Page {
	var pages []models.Page
	err := models.DbClient.MsClient.Where("id>? and id<?", 0, 10).Find(&page)
	if err != nil {
		return nil, err
	}
	return pages
}
