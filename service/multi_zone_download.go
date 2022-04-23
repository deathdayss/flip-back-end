package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary get a game
// @Description get a game, return a zip
// @Accept  plain
// @Produce  octet-stream
// @Param   game_id     header    int     true        "the game's id"
// @Success 200
// @Router /v1/download/game [GET]
func DownloadGameByZone(c *gin.Context) {
	gid := c.Query("game_id")
	game, err := repository.GetGameByZone(gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the game id is wrong",
		})
		return
	}
	if _, err := os.Stat("./storage/game/" + game.FileUrl); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the img name is wrong",
		})
		return
	}
	c.Writer.Header().Add("Content-Type", "Application/zip")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("inline;filename=%s", game.FileUrl))
	c.File("./storage/game/" + game.FileUrl)
}

// @Summary get a person's image
// @Description get a person's image
// @Accept  plain
// @Produce  png
// @Param   token     header    string     true        "token"
// @Success 200
// @Router /v1/download/personal [GET]
func DownloadPersonalByZone(c *gin.Context) {
	emailIt, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "unauth token",
		})
		return
	}
	email := emailIt.(string)
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
			"error":  "the img name is wrong",
		})
		return
	}
	contentType := ss[1]
	if _, err := os.Stat("./storage/personal/" + filename); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the img name is wrong",
		})
		return
	}
	if contentType == "jpg" || contentType == "jpeg" {
		c.Writer.Header().Add("Content-Type", "image/jpeg")
	} else {
		c.Writer.Header().Add("Content-Type", "image/png")
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("inline;filename=%s", filename))
	c.File("./storage/personal/" + filename)
}

// @Summary get a game's image
// @Description get a game's image
// @Accept  plain
// @Produce  png
// @Param   img_name     header    string     true        "the image name"
// @Success 200
// @Router /v1/download/img [GET]
func DownloadImgByZone(c *gin.Context) {
	filename := c.Query("img_name")
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "jpeg" && strings.ToLower(ss[1]) != "png") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the img name is wrong",
		})
		return
	}
	contentType := ss[1]
	if _, err := os.Stat("./storage/thumbnail/" + filename); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the img name is wrong",
		})
		return
	}
	if contentType == "jpg" || contentType == "jpeg" {
		c.Writer.Header().Add("Content-Type", "image/jpeg")
	} else {
		c.Writer.Header().Add("Content-Type", "image/png")
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("inline;filename=%s", filename))
	c.File("./storage/thumbnail/" + filename)
}
