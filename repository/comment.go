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

func UpComment(cid int, uid int) error {
	cu := &models.CommentUp{
		CID: cid,
		UID: uid,
	}
	if err := models.DbClient.MsClient.Model(cu).Create(cu).Error; err != nil {
		return err
	}
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

func GetCommentRanking(gid string, num int) (*[]models.Comment, error) {
	result := []models.Comment{}
	err := models.DbClient.MsClient.Where("g_id = ?", gid).
		Order("create_time DESC"). // order by create_time DESC
		Limit(num). // limit num
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