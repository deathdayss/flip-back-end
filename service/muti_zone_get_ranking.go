package service

import (
	"net/http"
	"strconv"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary get game rank by zone
// @Description get game rank by zone ordered by like, download or comment with default like
// @Accept  plain
// @Produce  json
// @Param   num     header    int     true        "num"
// @Param   zone     header    string     true        "zone"
// @Param   method     header    string     true        "like, download or comment with default like"
// @Success 200 {array} dto.RankItemByZone "{"status":200, "list":ranklist}"
// @Router /v1/rank/zone  [GET]
func GetGameRankingByZone(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	zone := c.Query("zone")
	if err != nil || len(zone) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "num or zone is wrong",
		})
		return
	}

	rankMtd, ok := c.GetQuery("method")
	if !ok || (rankMtd != "like" && rankMtd != "download" && rankMtd != "comment") {
		rankMtd = "time"
	}

	var offset int
	offsetStr, ok := c.GetQuery("offset")
	if !ok {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"status": 406,
				"error":  "offset if wrong",
			})
			return
		}
	}

	rankInfo, err := repository.GetGameRankingByZone(zone, num, offset, rankMtd)
	if err != nil || len(*rankInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}

	rankList := []dto.RankItemByZone{}
	for _, ri := range *rankInfo {
		rankList = append(rankList, dto.RankItemByZone{
			ID:          ri.ID,
			Name:        ri.Name,
			LikeNum:     ri.LikeNum,
			DownloadNum: ri.DownloadNum,
			CommentNum:  ri.CommentNum,
			Img:         ri.ImgUrl,
			AuthorName:  repository.FindNickName(ri.UID),
			ClickCount:  repository.GetGameClick(ri.ID),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List":   rankList,
	})
}

// @Summary search game
// @Description search game ordered by like, download or comment with default like
// @Accept  plain
// @Produce  json
// @Param   num     header    int     true        "num"
// @Param   keyword     header    string     true        "keyword"
// @Param   method     header    string     true        "like, download or comment with default like"
// @Success 200 {array} dto.RankItemByZone "{"status":200, "list":ranklist}"
// @Router /v1/search/game  [GET]
func SearchGameByZone(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	keyword := c.Query("keyword")
	if err != nil || len(keyword) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "keyword or num is missing",
		})
		return
	}
	rankMtd, ok := c.GetQuery("method")
	if !ok || (rankMtd != "like" && rankMtd != "download" && rankMtd != "comment") {
		rankMtd = "time"
	}
	var offset int
	offsetStr, ok := c.GetQuery("offset")
	if !ok {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"status": 406,
				"error":  "offset if wrong",
			})
			return
		}
	}
	rankInfo, err := repository.SearchGameByZone(keyword, num, offset, rankMtd)
	if err != nil || len(*rankInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}
	rankList := []dto.RankItemByZone{}
	for _, ri := range *rankInfo {
		rankList = append(rankList, dto.RankItemByZone{
			ID:          ri.ID,
			Name:        ri.Name,
			LikeNum:     ri.LikeNum,
			DownloadNum: ri.DownloadNum,
			CommentNum:  ri.CommentNum,
			Img:         ri.ImgUrl,
			AuthorName:  repository.FindNickName(ri.UID),
			ClickCount:  repository.GetGameClick(ri.ID),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List":   rankList,
	})
}
