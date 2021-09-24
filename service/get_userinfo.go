package service

import (
	"net/http"
	"strconv"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	id, ok1 := c.GetPostForm("id")
	if !ok1 || len(id) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "id is missing",
		})
		return
	}

	idnum, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "Invalid id",
		})
		return
	}
	if !repository.CheckID(idnum) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "id is not found",
		})
		return
	}

	userInfo, err := repository.FindUser(idnum)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no userinfo",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"userinfo": userInfo,
	})
}
