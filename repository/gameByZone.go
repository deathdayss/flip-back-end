package repository

import (
	"errors"

	"github.com/deathdayss/flip-back-end/models"
)

func GetGameRankingByZone(zone string, num int) (*[]models.GameRescale, error) {
	result := []models.GameRescale{}

	//sub_query := models.DbClient.MsClient.Model(&models.Zone{}).Select("id").Where("zone = ?", zone)

	err := models.DbClient.MsClient.Model(&models.GameRescale{}).
		Select("zones.id", "name", "like_num", "download_num", "comment_num", "img_url", "uid", "file_url").
		Joins("right join zones on zones.id = game_rescales.id").Where("zone = ?", zone).
		Order("like_num DESC"). // order by like_num DESC
		Limit(num).             // limit num
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
	//_ = sub_query //为了使语法不出错，没有任何用处的句子

	return &result, nil

}
