package service

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary allow user to like a game
// @Description allow user to like a game
// @Accept  plain
// @Produce  json
// @Param   gid     header    int     true        "GID"
// @Param   uid     header    int     true        "UID"
// @Success 200 {string} json  "{"status":http.StatusOK, "msg":"successfully like/unlike"}"
// @Router /v1/like/click  [GET]
func LikeOrUnlike(c *gin.Context) {
	gid := c.Query("GID")
	uid := c.Query("UID")
	err := repository.LikeOrUnlike(gid, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "can not like/unlike",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "successfully like/unlike",
		})
		return
	}
}

// @Summary get the like number of a game
// @Description get the like number of a game
// @Accept  plain
// @Produce  json
// @Param   gid     header    int     true        "GID"
// @Success 200 {int} count  "{"status":http.StatusOK, "count":count}"
// @Router /v1/like/num  [GET]
func GetLikeNum(c *gin.Context) {
	gid := c.Query("GID")
	count, err := repository.GetLikeNum(gid)
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

// @Summary check whether the user has liked the game
// @Description check whether the user has liked the game
// @Accept  plain
// @Produce  json
// @Param   gid     header    int     true        "GID"
// @Param   uid     header    int     true        "UID"
// @Success 200 {string} json  "{"status":http.StatusOK, "msg":true}"
// @Router /v1/like/check  [GET]
func HasLike(c *gin.Context) {
	gid := c.Query("GID")
	uid := c.Query("UID")
	hasLike := repository.IsExist(gid, uid, models.DbClient.MsClient)
	if hasLike {
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
