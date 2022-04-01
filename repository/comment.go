package repository

import (
	"errors"
	"strconv"
	"time"
	"gorm.io/gorm"
	"github.com/deathdayss/flip-back-end/models"
)

func AddComment(content string, gid, uid int) error {
	if !VerifyGame(strconv.Itoa(gid)) {
		return errors.New("no game whose gid is "+strconv.Itoa(gid))
	}
	if err := models.DbClient.MsClient.Model(&models.Game{ID: gid}).UpdateColumn("comment_num", gorm.Expr("comment_num + ?", 1)).Error; err != nil {
		return err
	}
	if err := models.DbClient.MsClient.Save(&models.Comment{
		Content: content,
		GID: gid,
		UID: uid,
		CreateTime: time.Now(),
		Up: 0,
		Down: 0,
	}).Error; err != nil {
		return errors.New("can not save the comment")
	}
	return nil
}

func UpComment(cid int) error {
	cm := models.Comment{
		ID: cid,
	}
	if err := models.DbClient.MsClient.Model(&cm).UpdateColumn("up", gorm.Expr("up + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func DownComment(cid int) error {
	cm := models.Comment{
		ID: cid,
	}
	if err := models.DbClient.MsClient.Model(&cm).UpdateColumn("down", gorm.Expr("down + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}