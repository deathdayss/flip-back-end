package service

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func LikeOrUnlike(c *gin.Context) {
	gid := c.Query("GID")
	uid := c.Query("UID")
	err := repository.LikeOrUnlike(gid, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "can not like/unlike",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg": "successfully like/unlike",
		})
		return
	}
}

func GetLikeNum(c *gin.Context) {
	gid := c.Query("GID")
	count, err := repository.GetLikeNum(gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"count": count,
		})
	}
}
func HasLike(c *gin.Context) {
	gid := c.Query("GID")
	uid := c.Query("UID")
	hasLike := repository.IsExist(gid, uid, models.DbClient.MsClient)
	if hasLike {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg": true,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg": false,
		})
		return
	}
}