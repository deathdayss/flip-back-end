package service

import (
	"net/http"

	//"github.com/deathdayss/flip-back-end/models"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func ClickOperation(c *gin.Context) {
	gid := c.Query("GID")
	uid := c.Query("UID")
	err := repository.ClickOperation(uid, gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "Can not click",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "finish click and update database",
		})
		return
	}
}
