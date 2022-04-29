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
func GetPersonalProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	method := c.Query("method")
	if err != nil || len(method) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "id or method is wrong",
		})
		return
	}
	rankInfo, err := repository.GetPersonalProduct(id, method)
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

func ChangeIcon(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	icon, ok2 := c.GetPostForm("icon")
	if !ok1 || !ok2 || len(icon) == 0 || len(email) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email or icon is missing",
		})
		return
	}
	pid, err := repository.Email2PID(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "can not auth the email",
		})
		return
	}

	if err := repository.ChangeIcon(pid, icon); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "can not set the new ico",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "new icon set successfully",
	})

}
