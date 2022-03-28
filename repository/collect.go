package repository

import (
	"strconv"
	"gorm.io/gorm"
	"github.com/deathdayss/flip-back-end/models"
)

func CollectOrUncollect(gid, uid string) error {
	tx := models.DbClient.MsClient.Begin()
	if IsCollectExist(gid, uid, tx) {
		collect := models.Collect{}
		tx.Where("user_id = ? AND game_id = ?", uid, gid).Take(&collect)
		tx.Delete(collect)
	} else {
		igid, err := strconv.Atoi(gid)
		if err != nil {
			tx.Rollback()
			return err
		}
		iuid, err := strconv.Atoi(uid)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Create(&models.Collect{
			GameID: igid,
			UserID: iuid,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func IsCollectExist(gid, uid string, tx *gorm.DB) bool {
	l := models.Collect{}
	if err := tx.Where("user_id = ? AND game_id = ?", uid, gid).First(&l).Error; err != nil {
		return false
	} else {
		return true
	}
}

func GetCollectNum(gid string) (int64, error) {
	var count int64
	err := models.DbClient.MsClient.Model(&models.Collect{}).Where("game_id = ?", gid).Count(&count).Error
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}

}
