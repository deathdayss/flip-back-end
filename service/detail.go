package service

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func GetUserDetail(c *gin.Context) {
	email, ok1 := c.GetQuery("email")
	password, ok2 := c.GetQuery("password")
	if !ok1 || !ok2 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email or password is missing",
		})
		return
	}
	if !repository.VerifyPerson(email, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "email or password is wrong",
		})
		return
	}
	detail, err := repository.GetUserDetail(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "the user has no details",
		})
		return
	}
	detailDto := dto.PersonDetail{
		Email: detail.Email,
		Age: detail.Age,
		Gender: detail.Gender,
		Birth: detail.Birth,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"detail": detailDto,
	})
}
var allowedField map[string]bool = map[string]bool{"Age":true, "Gender":true, "Birth":true}
func ChangeDetail(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	if !ok1 || !ok2 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email or password is missing",
		})
		return
	}
	if !repository.VerifyPerson(email, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "email or password is wrong",
		})
		return
	}
	fieldName, ok3 := c.GetPostForm("FieldKey")
	if !allowedField[fieldName] || !ok3 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "the field name is wrong",
		})
		return
	}
	fieldVal, ok4 := c.GetPostForm("FieldVal")
	if !ok4 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "the field value is missing",
		})
		return
	}
	if err := repository.ChangeDetail(email, fieldName, fieldVal); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "can not set the person",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg": "set successfully",
	})
}