package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func DownloadGame(c *gin.Context) {
	gid := c.Query("game_id")
	game, err := repository.GetGame(gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the game id is wrong",
		})
		return
	}
	if _, err := os.Stat("./storage/game/"+game.FileUrl); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	c.Writer.Header().Add("Content-Type", "Application/zip")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("inline;filename=%s", game.FileUrl))
	c.File("./storage/game/"+game.FileUrl)
}

func DownloadPersonal(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	if !ok1 || !ok2 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email or password is missing",
		})
		return
	}
	if !repository.VerifyPerson(email, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "email or password is wrong",
		})
		return
	}
	p, err := repository.FindPerson(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "email or password is wrong",
		})
		return
	}
	filename := repository.FindPersonal(p.ID)
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "jpeg" && strings.ToLower(ss[1]) != "png") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	contentType := ss[1]
	if _, err := os.Stat("./storage/personal/"+filename); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	if contentType == "jpg" || contentType == "jpeg" {
		c.Writer.Header().Add("Content-Type", "image/jpeg")
	} else {
		c.Writer.Header().Add("Content-Type", "image/png")
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("inline;filename=%s", filename))
	c.File("./storage/personal/"+filename)
}

func DownloadImg(c *gin.Context) {
	filename := c.Query("img_name")
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "jpeg" && strings.ToLower(ss[1]) != "png") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	contentType := ss[1]
	if _, err := os.Stat("./storage/thumbnail/"+filename); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	if contentType == "jpg" || contentType == "jpeg" {
		c.Writer.Header().Add("Content-Type", "image/jpeg")
	} else {
		c.Writer.Header().Add("Content-Type", "image/png")
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("inline;filename=%s", filename))
	c.File("./storage/thumbnail/"+filename)
}