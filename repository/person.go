package repository

import (
	"reflect"
	"strconv"

	"github.com/deathdayss/flip-back-end/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Email2PID(email string) (int, error) {
	p, err := FindPerson(email)
	if err != nil {
		return -1, err
	}
	return p.ID, nil
}
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
	saveName := "default.jpg"
	if fileType != "default" {
		saveName = strconv.Itoa(p.ID) + "." + fileType
	}
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
func FindPersonal(id int) string {
	pi := models.PersonImg{}
	if err := models.DbClient.MsClient.Where("uid=?", strconv.Itoa(id)).First(&pi).Error; err != nil {
		return "default.jpg"
	}
	if pi.URL == "" {
		return "default.jpg"
	}
	return pi.URL
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
func FindURL(id int) string {
	p := models.PersonImg{}
	if err := models.DbClient.MsClient.Where("id = ?", id).First(&p).Error; err != nil {
		return "default.jpg"
	} else {
		return p.URL
	}

}

func GetUserDetail(email string) (*models.PersonDetail, error) {
	detail := models.PersonDetail{Email: email}
	if err := models.DbClient.MsClient.Where("email=?", email).First(&detail).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &detail, nil
}

func ChangeDetail(email string, fieldName string, fieldVal string) error {
	tx := models.DbClient.MsClient.Begin()
	details := models.PersonDetail{Email: email}
	if err := tx.Clauses(clause.Locking{
		Strength: "UPDATE",
	}).Where("email=?", email).First(&details).Error; err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}
	if fieldName == "Age" {
		age, _ := strconv.Atoi(fieldVal)
		details.Age = age
	} else {
		pp := reflect.ValueOf(&details)
		field := pp.Elem().FieldByName(fieldName)
		field.SetString(fieldVal)
	}
	if err := tx.Where("email=?", email).Save(&details).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
