package repository

import (
	"errors"

	"github.com/deathdayss/flip-back-end/models"
)

func GetQuestion(num, offset int) (*[]models.Question, error) {
	result := []models.Question{}
	err := models.DbClient.MsClient.
		Limit(num). // limit num
		Offset(offset).
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

func ChangePassword(email, password string) bool {
	p := models.Person{}
	if err := models.DbClient.MsClient.Where("email = ?", string(email)).First(&p).Update("password", string(password)).Error; err != nil {
		return true
	} else {
		return false
	}
}

func FindAnswer(email string) (*models.Answer, error) {
	a := models.Answer{}
	err := models.DbClient.MsClient.Where("email = ?", string(email)).First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func CheckAnswer(email string) bool {
	_, err := FindAnswer(email)
	if err == nil {
		return true
	} else {
		return false
	}
}

func VerifyAnswer(email, answer string, question int) bool {
	a, err := FindAnswer(email)
	if err != nil {
		return false
	} else {
		if a.Question1 == question && a.Answer1 == answer || a.Question2 == question && a.Answer2 == answer || a.Question3 == question && a.Answer3 == answer {
			return true
		}
		return false
	}
}

func AddAnswer(email, answer1, answer2, answer3 string, question1, question2, question3 int) (string, error) {
	a := models.Answer{
		Email:     email,
		Question1: question1,
		Answer1:   answer1,
		Question2: question2,
		Answer2:   answer2,
		Question3: question3,
		Answer3:   answer3,
	}
	tx := models.DbClient.MsClient.Begin()
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return email, nil
}
