package service

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/deathdayss/flip-back-end/repository"
)

func UploadZip(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	gameID, ok3 := c.GetPostForm("game_id")
	if !ok1 || !ok2 || !ok3 || !repository.VerifyPerson(email, password) || !repository.VerifyGame(gameID) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"error" : "email or password or game is wrong",
		})
		return
	}
	file, err := c.FormFile("file_body")
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"status": http.StatusNoContent,
			"error" : "no data",
		})
		return
	}

	filename := file.Filename
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "zip" && strings.ToLower(ss[1]) != "ZIP") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	fileType := ss[1]
	saveName := gameID + "." + fileType
	if _, err := os.Stat("./storage/game/"+saveName); os.IsExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the game has been saved",
		})
		return
	}
	if err := c.SaveUploadedFile(file, "./storage/game/"+saveName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the game id is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg" : "the game has been saved successfully",
	})
}

func UploadImg(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	gameID, ok3 := c.GetPostForm("game_id")
	if !ok1 || !ok2 || !ok3 || !repository.VerifyPerson(email, password) || !repository.VerifyGame(gameID) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"error" : "email or password or game is wrong",
		})
		return
	}
	file, err := c.FormFile("file_body")
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"status": http.StatusNoContent,
			"error" : "no data",
		})
		return
	}

	filename := file.Filename
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "jpeg" && strings.ToLower(ss[1]) != "png") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	fileType := ss[1]
	saveName := gameID + "." + fileType
	if _, err := os.Stat("./storage/thumbnail/"+saveName); os.IsExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the game has been saved",
		})
		return
	}
	if err := c.SaveUploadedFile(file, "./storage/thumbnail/"+saveName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the game id is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg" : "the game thumbnail has been saved successfully",
	})

}