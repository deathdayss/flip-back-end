package service

import (
	"D/flip-back-end/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	if !ok1 || !ok2 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email or password is missing",
		})
		return
	}
	if repository.Verify(email, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "email or password is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successfully",
	})
}