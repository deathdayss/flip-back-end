package service

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary vertify a user's email and change password
// @Description change a user's password
// @Accept  plain
// @Produce  json
// @Param   email     body    string     true        "email"
// @Success 200 {json} string   "{"status":200, "message":email exists}"
// @Param   question     body    int     true        "question"
// @Param   answer     body    string     true        "answer"
// @Success 200 {json} string   "{"status":200, "message":vertify successfully}"
// @Param   newpwd     body    string     true        "newpwd"
// @Param   confirm     body    string     true        "confirm"
// @Success 200 {json} string   "{"status":200, "message":update successfully}"
// @Router /v1/notoken/change/password [POST]
func VertifyExist(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")

	if !ok1 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email is missing",
		})
		return
	}

	email = strings.TrimSpace(email)

	if !repository.CheckExistence(email) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "email is vertified",
	})
}

func VertifyAnswer(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	question, ok2 := c.GetPostForm("question")
	answer, ok3 := c.GetPostForm("answer")

	questionnum, err := strconv.Atoi(question)

	if !ok1 || !ok2 || !ok3 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email or answer is missing",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "Invalid id",
		})
		return
	}

	if !repository.VerifyAnswer(email, answer, questionnum) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "answer is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "answer is correct",
	})

}

func ChangePassword(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	newpwd, ok2 := c.GetPostForm("newpwd")
	confirm, ok3 := c.GetPostForm("confirm")
	if !ok1 || !ok2 || !ok3 || len(newpwd) == 0 || len(confirm) == 0 || len(email) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email, new password or confirm password is missing",
		})
		return
	}

	newpwd = strings.TrimSpace(newpwd)
	confirm = strings.TrimSpace(confirm)

	if newpwd != confirm {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "passswords do not match",
		})
		return
	}

	if newpwd == confirm {
		if !repository.ChangePassword(email, newpwd) {
			c.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "update successfully",
			})
		} else {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"status": 406,
				"error":  "can not update",
			})
			return
		}
	}

}
