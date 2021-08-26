package service

import (
	"D/flip-back-end/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	nickname, ok3 := c.GetPostForm("nickname")
	if !ok1 || !ok2 || !ok3 || len(email) == 0 || len(password) == 0 || len(nickname) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email, nickname or password is missing",
		})
		return
	}
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	nickname = strings.TrimSpace(nickname)
	if repository.CheckExistence(email) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email has been used",
		})
		return
	}
	if repository.Add(email, password, nickname) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "can not register",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "register successfully",
	})
}
