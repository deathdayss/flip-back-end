package repository

import "github.com/deathdayss/flip-back-end/models"

func GetProductInfo(GID int) (*models.ProductInfo, error) {
	p := models.ProductInfo{}
	if err := models.DbClient.MsClient.Where("ID = ?", GID).First(&p).Error; err != nil {
		return nil, err
	} else {
		return &p, nil
	}
}

func CheckGameExistence(GID int) bool {
	_, err := FindGame(GID)
	if err == nil {
		return true
	} else {
		return false
	}
}

func FindGame(GID int) (*models.ProductInfo, error) {
	g := models.ProductInfo{}
	err := models.DbClient.MsClient.Where("ID = ?", GID).First(&g).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}
