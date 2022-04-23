package service

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary get a user's detail
// @Description get a user's detail
// @Accept  plain
// @Produce  json
// @Param   token     header    string     true        "token"
// @Success 200 {object} dto.PersonDetail   "{"status":200, "detail":detail}"
// @Router /v1/user/change/detail [POST]
func GetUserDetail(c *gin.Context) {
	emailIt, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "unauth token",
		})
		return
	}
	email := emailIt.(string)
	detail, err := repository.GetUserDetail(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "the user has no details",
		})
		return
	}
	detailDto := dto.PersonDetail{
		Email:  detail.Email,
		Age:    detail.Age,
		Gender: detail.Gender,
		Birth:  detail.Birth,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"detail": detailDto,
	})
}

var allowedField map[string]bool = map[string]bool{"Age": true, "Gender": true, "Birth": true, "Sign": true, "Nickname": true}

// @Summary change a user's detail
// @Description change a user's detail
// @Accept  plain
// @Produce  json
// @Param   token     header    string     true        "token"
// @Param   FieldKey     body    string     true        "the attribute to be modified"
// @Param   FieldVal     body    string     true        "the attribute value to be modified"
// @Success 200 {json} string   "{"status":200, "msg": "set successfully"}"
// @Router /v1/user/detail [POST]
func ChangeDetail(c *gin.Context) {
	emailIt, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "unauth token",
		})
		return
	}
	email := emailIt.(string)
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
		"msg":    "set successfully",
	})
}
