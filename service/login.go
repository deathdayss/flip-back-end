package service

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/repository"
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
	if !repository.VerifyPerson(email, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "email or password is wrong",
		})
		return
	}
	token, err := generateToken(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "login successfully",
		"token":   token,
	})
}
