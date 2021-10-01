package repository

import "github.com/deathdayss/flip-back-end/models"

func GetProductInfo(GID int) (*models.Game, error) {
	g := models.Game{}
	if err := models.DbClient.MsClient.Where("ID = ?", GID).First(&g).Error; err != nil {
		return nil, err
	} else {
		return &g, nil
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

func FindGame(GID int) (*models.Game, error) {
	g := models.Game{}
	err := models.DbClient.MsClient.Where("ID = ?", GID).First(&g).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}

//以下都是用来返回productInfo列表中的单个变量的

func FindLikeNum(GID int) int {
	p := models.ProductInfo{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&p).Error; err != nil {
		return 0
	} else {
		return p.LikeNum
	}
}

func FindCollectionNum(GID int) int {
	p := models.ProductInfo{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&p).Error; err != nil {
		return 0
	} else {
		return p.CollectionNum
	}
}

func FindShareNum(GID int) int {
	p := models.ProductInfo{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&p).Error; err != nil {
		return 0
	} else {
		return p.ShareNum
	}
}

func FindClickCount(GID int) int {
	p := models.ProductInfo{}
	if err := models.DbClient.MsClient.Where("id = ?", GID).First(&p).Error; err != nil {
		return 0
	} else {
		return p.ClickCount
	}
}
