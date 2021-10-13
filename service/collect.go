package service

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func CollectOrUncollect(c *gin.Context) {
	gid := c.Query("GID")
	uid := c.Query("UID")
	err := repository.CollectOrUncollect(gid, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "can not collect/uncollect",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "successfully collect/uncollect",
		})
		return
	}
}

func GetCollectNum(c *gin.Context) {
	gid := c.Query("GID")
	count, err := repository.GetCollectNum(gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"count":  count,
		})
	}
}

func HasCollect(c *gin.Context) {
	gid := c.Query("GID")
	uid := c.Query("UID")
	hasCollect := repository.IsCollectExist(gid, uid, models.DbClient.MsClient)
	if hasCollect {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    true,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    false,
		})
		return
	}
}
