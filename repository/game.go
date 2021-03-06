package repository

import (
	"errors"
	"os"

	"github.com/deathdayss/flip-back-end/models"
)

func AddGame(name, email, imgUrl, fileUrl, zone string) (int, error) {
	author, err := FindPerson(email)
	if err != nil {
		return 0, err
	}
	uid := author.ID
	g := models.Game{
		Name:    name,
		ImgUrl:  imgUrl,
		UID:     uid,
		Zone:    zone,
		FileUrl: fileUrl,
	}
	if err := models.DbClient.MsClient.Create(&g).Error; err != nil {
		return 0, err
	}
	return g.ID, nil
}

func DeleteGame(id int) error {
	g := models.Game{
		ID: id,
	}
	DeleteFile(id)
	if err := models.DbClient.MsClient.Model(&models.Game{}).Delete(&g).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFile(id int) error {
	g := models.Game{
		ID: id,
	}
	if err := models.DbClient.MsClient.Model(&models.Game{}).Find(&g).Error; err != nil {
		return err
	}
	if g.FileUrl != "" {
		path := "./storage/game/" + g.FileUrl
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil

}
func UpdateGameFileUrl(id int, fileUrl string) error {
	g := models.Game{
		ID:      id,
		FileUrl: fileUrl,
	}
	if err := models.DbClient.MsClient.Model(&models.Game{}).Where("id=?", id).Save(&g).Error; err != nil {
		return err
	}
	return nil
}
func UpdateGameByID(id int, name, imgUrl, zone string, uid int) error {
	g := models.Game{
		ID:     id,
		Name:   name,
		ImgUrl: imgUrl,
		Zone:   zone,
		UID:    uid,
	}
	if err := models.DbClient.MsClient.Model(&models.Game{}).Where("id=?", id).Save(&g).Error; err != nil {
		return err
	}
	return nil
}

func VerifyGame(id string) bool {
	p := models.Game{}
	if err := models.DbClient.MsClient.Where("ID = ?", id).First(&p).Error; err != nil {
		return false
	} else {
		return true
	}
}

func GetGame(id string) (models.Game, error) {
	game := models.Game{}
	if err := models.DbClient.MsClient.Where("id = ?", id).First(&game).Error; err != nil {
		return game, err
	} else {
		return game, nil
	}
}

func GetGameRanking(zone string, num, offset int) (*[]models.Game, error) {
	result := []models.Game{}
	err := models.DbClient.MsClient.Where("zone = ?", zone).
		Order("like_num DESC"). // order by like_num DESC
		Limit(num).             // limit num
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

type PersonID struct {
	ID int `gorm:"column:id`
}

func SearchPerson(keyword string, num, offset int, mtd string) (*[]PersonID, error, int64) {
	result := make([]PersonID, 0)
	var order string
	switch mtd {
	case "like":
		order = "sum(like_num) DESC"
	case "download":
		order = "sum(download_num) DESC"
	case "comment":
		order = "sum(comment_num) DESC"
	default:
		order = "sum(like_num) DESC"
	}
	var recordNum int64 = 0
	if err := models.DbClient.MsClient.Debug().Model(&models.Person{}).Where("name LIKE ?", "%"+keyword+"%").Count(&recordNum).Error; err != nil {
		return nil, err, 0
	}
	err := models.DbClient.MsClient.Model(&models.Person{}).
		Select("people.id").
		Joins("JOIN games on people.id = games.uid").
		Where("people.nickname like ?", "%"+keyword+"%").
		Group("people.id").
		Order(order).Offset(offset).Limit(num).Find(&result).
		Error
	if err != nil {
		return nil, err, 0
	}
	return &result, nil, recordNum
}
func SearchGame(keyword string, num, offset int, mtd, zone string) (*[]models.Game, error, int64) {
	result := []models.Game{}
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
	// where := "1=1"
	// if zone != "" {
	// 	where = "zone="+zone
	// }
	var recordNum int64 = 0
	if err := models.DbClient.MsClient.Debug().Model(&models.Game{}).Where("name LIKE ?", "%"+keyword+"%").Count(&recordNum).Error; err != nil {
		return nil, err, 0
	}
	tx := models.DbClient.MsClient.Debug().Where("name LIKE ?", "%"+keyword+"%").
		Order(order).
		Limit(num).
		Offset(offset)
	var err error
	if zone != "" {
		err = tx.Where("zone=?", zone).Find(&result).Error
	} else {
		err = tx.Find(&result).Error
	}
	if err != nil {
		return nil, err, 0
	}
	actualLen := len(result)
	if actualLen == 0 {
		return nil, errors.New("No data"), 0
	}
	return &result, nil, recordNum
}

func GetGameRankingDownloading(zone string, num, offset int) (*[]models.Game, error) {
	result := []models.Game{}
	err := models.DbClient.MsClient.Where("zone = ?", zone).
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
