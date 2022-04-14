package service

import (
	"net/http"
	"strconv"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary get game according to the download number in the same zone
// @Description get game according to the download number in the same zone
// @Accept  plain
// @Produce  json
// @Param   num     header    int     true        "the number of the return itme"
// @Param   zone     header    int     true        "the zone"
// @Success 200 {array} dto.RankItem   "{"status":200, "List":list}"
// @Router /v1/rank/download [GET]
func GetGameRankingDownloading(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	zone := c.Query("zone")
	if err != nil || len(zone) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "num or zone is wrong",
		})
		return
	}
	var offset int
	offsetStr, ok := c.GetQuery("offset")
	if ! ok {
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
	rankInfo, err := repository.GetGameRankingDownloading(zone, num, offset)
	if err != nil || len(*rankInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}

	rankList := []dto.RankItem{}
	for _, ri := range *rankInfo {
		/*
			fp, err := ioutil.ReadFile("./storage/thumbnail/"+ri.ImgUrl)
			if err != nil {
				fp, _ = ioutil.ReadFile("./storage/thumbnail/not_found.png")
			}*/
		rankList = append(rankList, dto.RankItem{
			ID:          ri.ID,
			Name:        ri.Name,
			LikeNum:     ri.LikeNum,
			DownloadNum: ri.DownloadNum,
			CommentNum:  ri.CommentNum,
			Img:         ri.ImgUrl,
			AuthorName:  repository.FindNickName(ri.UID),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List":   rankList,
	})
}
