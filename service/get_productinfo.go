package service

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func GetProductInfo(c *gin.Context) {
	id, ok1 := c.GetPostForm("id")
	if !ok1 || len(id) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "id is missing",
		})
		return
	}

	if !repository.VerifyGame(id) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "id is not found",
		})
		return
	}

	userInfo, err := repository.GetGame(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no productinfo",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"userinfo": userInfo,
	})
}
