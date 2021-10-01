package repository

import (
	"strconv"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/jinzhu/gorm"
)

func ShareOperation(gid string) error {
	//对share表进行更新
	ts := models.DbClient.MsClient.Begin()
	if IsGamePersonExist(gid, ts) {
		share := models.Share{}
		ts.Where("game_id = ?", gid).Take(&share)
		share.ShareNum = share.ShareNum + 1
		ts.Update(share)

	} else {
		igid, err := strconv.Atoi(gid) //gid不合法
		if err != nil {
			ts.Rollback()
			return err
		}
		err = ts.Create(&models.Share{ //如果用户以前没有分享过，就给他创建一个分享数据，分享一次
			GameID:   igid,
			ShareNum: 1,
		}).Error
		if err != nil {
			ts.Rollback()
			return err
		}
	}
	ts.Commit()

	//对productInfo进行更新
	ti := models.DbClient.MsClient.Begin()
	shareRecord := models.ProductInfo{}
	ti.Where("game_id = ?", gid).Take(&shareRecord)
	shareRecord.ShareNum = shareRecord.ShareNum + 1
	ti.Update(shareRecord)
	ti.Commit()

	return nil
}

func IsGamePersonExist(gid string, tx *gorm.DB) bool {
	l := models.Share{}
	if err := tx.Where("game_id = ? ", gid).First(&l).Error; err != nil {
		return false
	} else {
		return true
	}
}

func GetShareNum(gid string) (int, error) {
	s := models.Share{}

	err := models.DbClient.MsClient.Model(&models.Like{}).Where("game_id = ?", gid).First(&s).Error
	if err != nil {
		return 0, err
	} else {
		return s.ShareNum, nil
	}
}
