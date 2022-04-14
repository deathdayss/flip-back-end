package repository

import (
	"errors"

	"github.com/deathdayss/flip-back-end/models"
)

func GetAuthorRankingByZone(zone string, num int) (*[]models.Author, error) {
	result := []models.Author{}

	err := models.DbClient.MsClient.Debug().Model(&models.Game{}).
		Select("url", "person.nickname as nickname", "sum(games.like_num) as like_num").
		Joins("left join person on games.uid = person.id").
		Joins("join person_imgs on games.uid = person_imgs.uid").Where("zone=?", zone).
		Group("games.uid").
		Order("sum(like_num)").
		Limit(num).
		Find(&result).Error

	if err != nil {
		return nil, err
	}
	actualLen := len(result)
	if actualLen == 0 {
		return nil, errors.New("No data")
	}
	if actualLen < num {
		for i := actualLen; i < num; i++ {
			result = append(result, result[i%actualLen])
		}
	}

	return &result, nil

}
