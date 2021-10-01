package repository

import (
	"time"

	"github.com/deathdayss/flip-back-end/models"
)

func AddGame(name, email, imgUrl, zone string) error {
	author, err := FindPerson(email)
	if err != nil {
		return err
	}
	uid := author.ID
	g := models.Game{
		Name:   name,
		ImgUrl: imgUrl,
		UID:    uid,
		Zone:   zone,
	}
	if err := models.DbClient.MsClient.Create(&g).Error; err != nil {
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

func GetGameRanking(zone string, num int) (*[]models.ProductInfo, error) {
	result := []models.ProductInfo{}
	err := models.DbClient.MsClient.Where("zone = ?", zone).
		Order("like_num DESC"). // order by like_num DESC
		Limit(num).             // limit num
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	actualLen := len(result)
	if actualLen < num {
		for i := actualLen; i < num; i++ {
			result = append(result, result[i%actualLen])
		}
	}
	return &result, nil
}

func GetGameRankingDownloading(zone string, num int) (*[]models.ProductInfo, error) {
	result := []models.ProductInfo{}
	err := models.DbClient.MsClient.Where("zone = ?", zone).
		Order("download_num DESC").
		Limit(num).
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	actualLen := len(result)
	if actualLen < num {
		for i := actualLen; i < num; i++ {
			result = append(result, result[i%actualLen])
		}
	}
	return &result, nil
}

//以下都是用来返回game列表中的参数的

func FindGameName(GID int) string {
	g := models.Game{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&g).Error; err != nil {
		return "Undefined"
	} else {
		return g.Name
	}
}

func FindGameIntroduction(GID int) string {
	g := models.Game{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&g).Error; err != nil {
		return "Undefined"
	} else {
		return g.Introduction
	}
}

func FindGameImgUrl(GID int) string {
	g := models.Game{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&g).Error; err != nil {
		return "Undefined"
	} else {
		return g.ImgUrl
	}
}

func FindGameFileUrl(GID int) string {
	g := models.Game{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&g).Error; err != nil {
		return "Undefined"
	} else {
		return g.FileUrl
	}
}

func FindGameCreateAt(GID int) (time.Time, error) {
	g := models.Game{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&g).Error; err != nil {
		return time.Now(), err
	} else {
		return g.CreateAt, nil
	}
}

func FindUserID(GID int) int {
	g := models.Game{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&g).Error; err != nil {
		return 0
	} else {
		return g.UID
	}
}

/*
调整数据库结构之前的样子
func GetGameRanking(zone string, num int) (*[]models.Game, error) {
	result := []models.Game{}
	err := models.DbClient.MsClient.Where("zone = ?", zone).
		Order("like_num DESC"). // order by like_num DESC
		Limit(num).             // limit num
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	actualLen := len(result)
	if actualLen < num {
		for i := actualLen; i < num; i++ {
			result = append(result, result[i%actualLen])
		}
	}
	return &result, nil
}

func GetGameRankingDownloading(zone string, num int) (*[]models.Game, error) {
	result := []models.Game{}
	err := models.DbClient.MsClient.Where("zone = ?", zone).
		Order("download_num DESC").
		Limit(num).
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	actualLen := len(result)
	if actualLen < num {
		for i := actualLen; i < num; i++ {
			result = append(result, result[i%actualLen])
		}
	}
	return &result, nil
}
*/
