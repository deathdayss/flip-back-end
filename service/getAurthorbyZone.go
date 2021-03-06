package service

import (
	"net/http"
	"strconv"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary get anuthor rank by zone
// @Description get anuthor rank by zone
// @Accept  plain
// @Produce  json
// @Param   num     header    int     true        "num"
// @Param   zone     header    string     true        "zone"
// @Success 200 {array} dto.AuthorItem "{"status":200, "list":ranklist}"
// @Router /v1/rank/author  [GET]
func GetAurthorRankingByZone(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	zone := c.Query("zone")
	if err != nil || len(zone) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "num or zone is wrong",
		})
		return
	}
	rankInfo, err := repository.GetAuthorRankingByZone(zone, num)
	if err != nil || len(*rankInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}

	rankList := []dto.AuthorItem{}
	for _, ri := range *rankInfo {
		rankList = append(rankList, dto.AuthorItem{
			NickName: ri.NickName,
			LikeNum:  ri.LikeNum,
			URL:      ri.URL,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List":   rankList,
	})
}
