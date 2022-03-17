package repository

import (
	"strconv"

	"github.com/deathdayss/flip-back-end/models"
)

func CheckExistence(email string) bool {
	_, err := FindPerson(email)
	if err == nil {
		return true
	} else {
		return false
	}
}

func FindPerson(email string) (*models.Person, error) {
	p := models.Person{}
	err := models.DbClient.MsClient.Where("email = ?", string(email)).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func AddUser(email, password, nickname string, fileType string) (string, error) {
	p := models.Person{
		Email:    email,
		Password: password,
		Nickname: nickname,
	}
	tx := models.DbClient.MsClient.Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	saveName := strconv.Itoa(p.ID) + "." + fileType
	pi := models.PersonImg{
		UID: p.ID,
		URL: saveName,
	}
	if err := tx.Create(&pi).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return saveName, nil
}

func VerifyPerson(email, password string) bool {
	p := models.Person{}
	if err := models.DbClient.MsClient.Where("email = ? AND password = ?", string(email), string(password)).First(&p).Error; err != nil {
		return false
	} else {
		return true
	}
}

func FindNickName(id int) string {
	p := models.Person{}
	if err := models.DbClient.MsClient.Where("id = ?", id).First(&p).Error; err != nil {
		return "anonymity"
	} else {
		return p.Nickname
	}
}
