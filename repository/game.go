package repository

import "github.com/deathdayss/flip-back-end/models"

func AddGame(name, email, imgUrl, fileUrl, zone string) (int, error) {
	author, err := FindPerson(email)
	if err != nil {
		return 0, err
	}
	uid := author.ID
	g := models.Game{
		Name:   name,
		ImgUrl: imgUrl,
		UID:    uid,
		Zone:   zone,
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
	if err := models.DbClient.MsClient.Model(&models.Game{}).Delete(&g).Error; err != nil {
		return err
	}
	return nil
}

func UpdateGameByID(id int, name, imgUrl, zone string, uid int) error {
	g := models.Game{
		ID: id,
		Name: name,
		ImgUrl: imgUrl,
		Zone: zone,
		UID: uid,
	}
	if err := models.DbClient.MsClient.Model(&models.Game{}).Update(&g).Error; err != nil {
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

func GetGameRanking(zone string, num int) (*[]models.Game, error) {
	result := []models.Game{}
	err := models.DbClient.MsClient.Where("zone = ?", zone).
		Order("like_num DESC"). // order by like_num DESC
		Limit(num). // limit num
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
