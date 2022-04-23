package service

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

var ProcessMapByZone sync.Map = sync.Map{}

func UploadZipByZone(c *gin.Context) {
	emailIt, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "unauth token",
		})
		return
	}
	email := emailIt.(string)
	file, err := c.FormFile("file_body")
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"status": http.StatusNoContent,
			"error":  "no data",
		})
		return
	}

	filename := file.Filename
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "zip" && strings.ToLower(ss[1]) != "ZIP") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the img name is wrong",
		})
		return
	}
	fileType := ss[1]
	//此处暂时不用处理zone, zone在另一个表中
	//processID应该就是GID
	processID, err := repository.AddGameByZone("", email, "", "")
	saveName := strconv.Itoa(processID) + "." + fileType
	if _, err := os.Stat("./storage/game/" + saveName); os.IsExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the game has been saved",
		})
		return
	}
	if err := c.SaveUploadedFile(file, "./storage/game/"+saveName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the game id is wrong",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "can not create game",
		})
		return
	}

	repository.UpdateGameFileUrlByZone(processID, saveName)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"ID":     processID,
	})

	flagChan := make(chan bool)
	go func() {
		ProcessMapByZone.Store(processID, flagChan)
		if timeoutByZone(flagChan) {
			repository.DeleteGame(processID)
		}
		ProcessMapByZone.Delete(processID)
	}()
}

func timeoutByZone(c chan bool) bool {
	time.AfterFunc(5*time.Minute, func() {
		c <- true
	})
	r := <-c
	return r
}

func UploadInfoByZone(c *gin.Context) {
	emailIt, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "unauth token",
		})
		return
	}
	email := emailIt.(string)
	gameID, ok3 := c.GetPostForm("game_id")
	gameName, ok4 := c.GetPostForm("game_name")
	zone, ok5 := c.GetPostForm("zone")

	if !ok3 || !ok4 || !ok5 || !repository.VerifyGame(gameID) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"error":  "email or password or game is wrong",
		})
		return
	}
	file, err := c.FormFile("file_body")
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"status": http.StatusNoContent,
			"error":  "no data",
		})
		return
	}

	filename := file.Filename
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "jpeg" && strings.ToLower(ss[1]) != "png") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the img name is wrong",
		})
		return
	}
	fileType := ss[1]
	saveName := gameID + "." + fileType
	if _, err := os.Stat("./storage/thumbnail/" + saveName); os.IsExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the game has been saved",
		})
		return
	}
	if err := c.SaveUploadedFile(file, "./storage/thumbnail/"+saveName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the game id is wrong",
		})
		return
	}
	user, err := repository.FindPerson(email)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"error":  "email or password or game is wrong",
		})
		return
	}
	uid := user.ID
	iGID, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the game id is wrong",
		})
		return
	}
	msgChan, ok := ProcessMap.Load(iGID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the game can not update",
		})
		return
	}
	err = repository.UpdateGameByIDByZone(iGID, gameName, saveName, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the game can not update",
		})
		return
	}
	err = repository.UpdateZoneByZone(iGID, zone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "the zone is wrong",
		})
		return
	}

	m, _ := msgChan.(chan bool)
	m <- false
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "the game thumbnail has been saved successfully",
	})

}
