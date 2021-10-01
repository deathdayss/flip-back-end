package service

import (
	"net/http"

	//"github.com/deathdayss/flip-back-end/models"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func ShareOperation(c *gin.Context) {
	gid := c.Query("GID")
	err := repository.ShareOperation(gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "can not share",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "successfully share",
		})
		return
	}
}
