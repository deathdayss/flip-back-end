package repository

import (
	"errors"
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
	if fieldName == "Sign" {
		sign := models.Signature{Email: email, Content: fieldVal}
		if err := tx.Model(sign).Save(sign).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}
	if fieldName == "Nickname" {
		if err := tx.Model(&models.Person{}).Where("email=?", email).Update("nickname", fieldVal).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}
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

func ChangeIcon(id int, fileType string) (string, error) {

	saveName := strconv.Itoa(id) + "." + fileType
	tx := models.DbClient.MsClient.Begin()

	if err := tx.Model(&models.PersonImg{}).Where("id=?", id).Update("url", saveName).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return saveName, nil

}

func GetPersonalProduct(id int, mtd string) (*[]models.GameRescale, error) {
	result := []models.GameRescale{}

	//sub_query := models.DbClient.MsClient.Model(&models.Zone{}).Select("id").Where("zone = ?", zone)

	var order string
	switch mtd {
	case "like":
		order = "like_num DESC"
	case "download":
		order = "download_num DESC"
	case "comment":
		order = "comment_num DESC"
	default:
		order = "like_num DESC"
	}

	err := models.DbClient.MsClient.Debug().Model(&models.GameRescale{}).
		Select("id", "name", "like_num", "download_num", "comment_num", "img_url", "uid", "file_url").
		Where("uid = ?", id).
		Order(order).
		Find(&result).Error

	if err != nil {
		return nil, err
	}
	actualLen := len(result)
	if actualLen == 0 {
		return nil, errors.New("No data")
	}
	return &result, nil

}
