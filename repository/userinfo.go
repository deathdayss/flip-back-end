package repository

import "github.com/deathdayss/flip-back-end/models"

func CheckID(ID int) bool {
	_, err := FindUser(ID)
	if err == nil {
		return true
	} else {
		return false
	}
}

func FindUser(ID int) (*models.Person, error) {
	p := models.Person{}
	err := models.DbClient.MsClient.Where("ID = ?", int(ID)).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}
