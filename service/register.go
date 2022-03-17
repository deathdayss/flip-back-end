package service

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	email, ok1 := c.GetPostForm("email")
	password, ok2 := c.GetPostForm("password")
	nickname, ok3 := c.GetPostForm("nickname")
	if !ok1 || !ok2 || !ok3 || len(email) == 0 || len(password) == 0 || len(nickname) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email, nickname or password is missing",
		})
		return
	}
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	nickname = strings.TrimSpace(nickname)
	if repository.CheckExistence(email) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email has been used",
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
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "png") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	fileType := ss[1]
	pid, err := repository.AddUser(email, password, nickname)
	saveName := strconv.Itoa(pid) + "." + fileType
	if _, err := os.Stat("./storage/personal/"+saveName); os.IsExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the user has been saved",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "can not register",
		})
		return
	}
	if err := c.SaveUploadedFile(file, "./storage/personal/"+saveName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the person id is wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "register successfully",
	})
}
