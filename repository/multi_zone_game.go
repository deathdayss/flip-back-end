package repository

import (
	"errors"
	"strings"

	"github.com/deathdayss/flip-back-end/models"
)

func AddGameByZone(name, email, imgUrl, fileUrl string) (int, error) {
	author, err := FindPerson(email)
	if err != nil {
		return 0, err
	}
	uid := author.ID
	g := models.GameRescale{ //这里替换了
		Name:    name,
		ImgUrl:  imgUrl,
		UID:     uid,
		FileUrl: fileUrl,
	}

	if err := models.DbClient.MsClient.Create(&g).Error; err != nil {
		return 0, err
	}

	return g.ID, nil
}

func UpdateGameFileUrlByZone(id int, fileUrl string) error {
	g := models.GameRescale{
		ID:      id,
		FileUrl: fileUrl,
	}
	if err := models.DbClient.MsClient.Model(&models.Game{}).Save(&g).Error; err != nil {
		return err
	}
	return nil
}

func UpdateGameByIDByZone(id int, name, imgUrl string, uid int) error {
	g := models.Game{
		ID:     id,
		Name:   name,
		ImgUrl: imgUrl,
		UID:    uid,
	}
	if err := models.DbClient.MsClient.Model(&models.Game{}).Save(&g).Error; err != nil {
		return err
	}
	return nil
}

func UpdateZoneByZone(id int, zone string) error {
	context := strings.Fields(zone)
	for i := range context {
		temp := models.Zone{
			ID:   id,
			Zone: context[i],
		}

		if err := models.DbClient.MsClient.Model(&models.Zone{}).Save(&temp).Error; err != nil {
			return err
		}
	}

	return nil
}

func VerifyGameByZone(id string) bool {
	p := models.GameRescale{}
	if err := models.DbClient.MsClient.Where("ID = ?", id).First(&p).Error; err != nil {
		return false
	} else {
		return true
	}
}

func GetGameByZone(id string) (models.GameRescale, error) {
	game := models.GameRescale{}
	if err := models.DbClient.MsClient.Where("id = ?", id).First(&game).Error; err != nil {
		return game, err
	} else {
		return game, nil
	}
}

func SearchGameByZone(keyword string, num, offset int, mtd string) (*[]models.GameRescale, error) {
	result := []models.GameRescale{}
	var order string
	switch mtd {
	case "like":
		order = "like_num DESC"
	case "download":
		order = "download_num DESC"
	case "comment":
		order = "comment_num DESC"
	default:
		order = "like_num DESC"
	}

	err := models.DbClient.MsClient.Model(&models.GameRescale{}).
		Select("zones.id", "name", "like_num", "download_num", "comment_num", "img_url", "uid", "file_url").
		Joins("right join zones on zones.id = game_rescales.id").Where("name LIKE ?", "%"+keyword+"%").
		Order(order).
		Limit(num).
		Offset(offset).
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

// It provides 3 kinds of ranking method, by like, by
func GetGameRankingByZone(zone string, num int, offset int, mtd string) (*[]models.GameRescale, error) {
	result := []models.GameRescale{}

	//sub_query := models.DbClient.MsClient.Model(&models.Zone{}).Select("id").Where("zone = ?", zone)

	var order string
	switch mtd {
	case "like":
		order = "like_num DESC"
	case "download":
		order = "download_num DESC"
	case "comment":
		order = "comment_num DESC"
	default:
		order = "like_num DESC"
	}

	err := models.DbClient.MsClient.Model(&models.GameRescale{}).
		Select("zones.id", "name", "like_num", "download_num", "comment_num", "img_url", "uid", "file_url").
		Joins("right join zones on zones.id = game_rescales.id").Where("zone = ?", zone).
		Order(order).
		Limit(num).
		Offset(offset).
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

func GetGameRankingDownloadingByZone(zone string, num, offset int) (*[]models.GameRescale, error) {
	result := []models.GameRescale{}
	err := models.DbClient.MsClient.Model(&models.GameRescale{}).
		Select("zones.id", "name", "like_num", "download_num", "comment_num", "img_url", "uid", "file_url").
		Joins("right join zones on zones.id = game_rescales.id").Where("zone = ?", zone).
		Order("download_num DESC").
		Limit(num).
		Offset(offset).
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

func GetGameClick(id int) int {
	c := models.Click{}
	if err := models.DbClient.MsClient.Where("game_id = ?", id).First(&c).Error; err != nil {
		return 0
	} else {
		return c.Count
	}
}
