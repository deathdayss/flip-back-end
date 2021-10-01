package service

import (
	"net/http"
	//"io/ioutil"
	"strconv"

	"github.com/deathdayss/flip-back-end/repository"

	//"github.com/deathdayss/flip-back-end/dto"
	"github.com/gin-gonic/gin"
)

func GetProductInfo(c *gin.Context) { //gin.Context用于处理http请求
	GID, err := strconv.Atoi(c.Query("GID")) //获取请求的pid，把String转为int
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "The GID is illegal",
		})
		return
	}

	if !repository.CheckGameExistence(GID) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "No such game",
		})
		return
	}

	//上面检查输入是否合法完毕
	//下面返回查询到的值

	game, err := repository.GetProductInfo(GID)
	c.JSON(http.StatusOK, gin.H{
		"Status":         200,
		"ID":             game.ID,
		"game_name":      game.Name,
		"like_num":       repository.FindLikeNum(GID),
		"collection_num": repository.FindCollectionNum(GID),
		"share_num":      repository.FindShareNum(GID),
		"introduction":   game.Introduction,
		"img_url":        game.ImgUrl,
		"uid":            game.UID,
		"time":           game.CreateAt,
		"click_count":    repository.FindClickCount(GID),
	})
}
