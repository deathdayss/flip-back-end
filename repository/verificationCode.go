package repository

import "github.com/deathdayss/flip-back-end/models"

func GetCode(ID int) (*models.Code, error) {
	p := models.Code{}
	if err := models.DbClient.MsClient.Where("ID = ?", ID).First(&p).Error; err != nil {
		return nil, err
	} else {
		return &p, nil
	}
}
