package repository

import "github.com/deathdayss/flip-back-end/models"

func GetProductInfo(pid int) (*models.Game, error) {
	result := models.Game{}

	return &result, nil
}

func CheckGameExistence(pid int) bool {
	_, err := FindGame(pid)
	if err == nil {
		return true
	} else {
		return false
	}
}

func FindGame(pid int) (*models.Game, error) {
	g := models.Game{}
	err := models.DbClient.MsClient.Where("ID = ?", pid).First(&g).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}
