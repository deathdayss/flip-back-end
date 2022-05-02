package service

import (
	"net/http"
	"os"
	"strconv"
	"strings"

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

	if !ok1 || len(email) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email is missing",
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

	url := repository.FindURL(pid)

	if url != "default.jpg" {

		if err := os.Remove("./storage/personal/" + url); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "the person id is wrong",
			})
			return
		}
	}

	file, err := c.FormFile("file_body")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "can not auth the file",
		})
		return
	}

	filename := file.Filename
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "png") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the img name is wrong",
		})
		return
	}
	fileType := ss[1]

	saveName, err := repository.ChangeIcon(pid, fileType)

	if _, err := os.Stat("./storage/personal/" + saveName); os.IsExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the user has been saved",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "can not upload",
		})
		return
	}

	if err := c.SaveUploadedFile(file, "./storage/personal/"+saveName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the person id is wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Change Icon successfully",
	})

}
