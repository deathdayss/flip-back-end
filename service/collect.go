package service

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary allow user to collect a game
// @Description allow user to collect a game
// @Accept  plain
// @Produce  json
// @Param   gid     header    int     true        "GID"
// @Param   uid     header    int     true        "UID"
// @Success 200 {string} json  "{"status":http.StatusOK, "msg":"successfully collect/uncollect"}"
// @Router /v1/collect/click  [GET]
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

// @Summary get the collect number of a game
// @Description get the collect number of a game
// @Accept  plain
// @Produce  json
// @Param   gid     header    int     true        "GID"
// @Success 200 {int} count  "{"status":http.StatusOK, "count":count}"
// @Router /v1/collect/num  [GET]
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

// @Summary check whether the user has collected the game
// @Description check whether the user has collected the game
// @Accept  plain
// @Produce  json
// @Param   gid     header    int     true        "GID"
// @Param   uid     header    int     true        "UID"
// @Success 200 {string} json  "{"status":http.StatusOK, "msg":true}"
// @Router /v1/collect/check  [GET]
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
