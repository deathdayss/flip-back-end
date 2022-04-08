package service

import (
	"net/http"
	"strings"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {

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
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "email exists",
		})
	}

	question, ok2 := c.GetPostForm("question")
	answer, ok3 := c.GetPostForm("answer")

	if !ok2 || !ok3 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "answer is missing",
		})
		return
	}

	if !repository.VerifyAnswer(email, question, answer) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "answer is wrong",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "vertify successfully",
		})
	}

	newpwd, ok4 := c.GetPostForm("newpwd")
	confirm, ok5 := c.GetPostForm("confirm")
	if !ok4 || !ok5 || len(newpwd) == 0 || len(confirm) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "new password or confirm password is missing",
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
