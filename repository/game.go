package repository

import "github.com/deathdayss/flip-back-end/models"

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

func GetGameRanking(zone string, num int) (*[]models.Game, error) {
	result := []models.Game{}
	err := models.DbClient.MsClient.Where("zone = ?", zone).
		Order("like_num DESC").
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
