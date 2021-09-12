package repository

import "github.com/deathdayss/flip-back-end/models"

func GetProductInfo(pid int) (*models.ProductInfo, error) {
	p := models.ProductInfo{}
	if err := models.DbClient.MsClient.Where("ID = ?", pid).First(&p).Error; err != nil {
		return nil, err
	} else {
		return &p, nil
	}
}

func CheckGameExistence(pid int) bool {
	_, err := FindGame(pid)
	if err == nil {
		return true
	} else {
		return false
	}
}

func FindGame(pid int) (*models.ProductInfo, error) {
	g := models.ProductInfo{}
	err := models.DbClient.MsClient.Where("ID = ?", pid).First(&g).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}
