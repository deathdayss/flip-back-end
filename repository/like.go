package repository

import (
	"strconv"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/jinzhu/gorm"
)

func LikeOrUnlike(gid, uid string) error {
	tx := models.DbClient.MsClient.Begin()
	if IsExist(gid, uid, tx) {
		like := models.Like{}
		tx.Where("user_id = ? AND game_id = ?", uid, gid).Take(&like)
		tx.Delete(like)
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
		err = tx.Create(&models.Like{
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

func IsExist(gid, uid string, tx *gorm.DB) bool {
	l := models.Like{}
	if err := tx.Where("user_id = ? AND game_id = ?", uid, gid).First(&l).Error; err != nil {
		return false
	} else {
		return true
	}
}

func GetLikeNum(gid string) (int, error) {
	var count int
	err := models.DbClient.MsClient.Model(&models.Like{}).Where("game_id = ?", gid).Count(&count).Error
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}
