package repository

import (
	"strconv"

	"time"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/jinzhu/gorm"
)

func ClickOperation(gid, uid string) error {
	tc := models.DbClient.MsClient.Begin()
	if IsExist(gid, uid, tc) {
		click := models.Click{}
		date := time.Now()

		tc.Where("user_id = ? AND game_id = ?", uid, gid).Take(&click)
		if CompareDate(date, click.Date) { // 如果这个日期已经出现过，那么就不可以点击+1
			return nil
		} else { //这个日期没出现过，点击+1，同时更新数click表中的日期记录操作
			click.ClickCount = click.ClickCount + 1
			click.Date = date
		}
		tc.Update(click)

	} else {
		igid, err := strconv.Atoi(gid)
		if err != nil {
			tc.Rollback()
			return err
		}
		iuid, err := strconv.Atoi(uid)
		if err != nil {
			tc.Rollback()
			return err
		}
		date := time.Now()
		err = tc.Create(&models.Click{
			UserID:     iuid,
			GameID:     igid,
			Date:       date,
			ClickCount: 1,
		}).Error
		if err != nil {
			tc.Rollback()
			return err
		}
	}
	tc.Commit()
	return nil
}

func CompareDate(date1, date2 time.Time) bool {
	if date1.Year() == date2.Year() {
		if date1.Month() == date2.Month() {
			if date1.Day() == date2.Day() {
				return true
			}
		}
	}
	return false
}

func IsClickExist(gid, uid string, tx *gorm.DB) bool {
	l := models.Click{}
	if err := tx.Where("user_id = ? AND game_id = ?", uid, gid).First(&l).Error; err != nil {
		return false
	} else {
		return true
	}
}

func GetClickCount(gid string) (int, error) {
	c := models.Click{}
	err := models.DbClient.MsClient.Model(&models.Like{}).Where("game_id = ?", gid).First(&c).Error
	if err != nil {
		return 0, err
	} else {
		return c.ClickCount, nil
	}
}
