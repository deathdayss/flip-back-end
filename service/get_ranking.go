package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func GetGameRanking(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	zone := c.Query("zone")
	if err != nil || len(zone) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "num or zone is wrong",
		})
		return
	}
	rankInfo, err := repository.GetGameRanking(zone, num)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}

	rankList := []dto.RankItem{}
	for _, ri := range (*rankInfo) {
		fp, err := os.ReadFile("./storage/thumbnail/"+ri.ImgUrl)
		if err != nil {
			fp, _ = os.ReadFile("./storage/thumbnail/not_found.png")
		}
		rankList = append(rankList, dto.RankItem{
			ID : ri.ID,
			Name : ri.Name,
			LikeNum: ri.LikeNum,
			DownloadNum : ri.DownloadNum,
			CommentNum : ri.CommentNum,
			Img : fp,
			AuthorName : repository.FindNickName(ri.UID),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List": rankList,
	})
}