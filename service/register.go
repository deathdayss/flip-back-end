package service

import (
	"net/http"
	"os"
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

	question1, ok4 := c.GetPostForm("question1")
	answer1, ok5 := c.GetPostForm("answer1")
	question2, ok6 := c.GetPostForm("question2")
	answer2, ok7 := c.GetPostForm("answer2")
	question3, ok8 := c.GetPostForm("question3")
	answer3, ok9 := c.GetPostForm("answer3")

	if !ok4 || !ok5 || !ok6 || !ok7 || !ok8 || !ok9 || len(answer1) == 0 || len(answer2) == 0 || len(answer3) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "answer missing",
		})
		return
	}

	if !repository.CheckAnswer(email) {
		_, err := repository.AddAnswer(email, question1, answer1, question2, answer2, question3, answer3)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "cannot save answer",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "add answer successfully",
			})
		}
	}

	file, err := c.FormFile("file_body")
	var fileType string = "default"
	if err == nil {
		filename := file.Filename
		ss := strings.Split(filename, ".")
		if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "png") {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "the img name is wrong",
			})
			return
		}
		fileType = ss[1]
	}
	saveName, err := repository.AddUser(email, password, nickname, fileType)
	if saveName != "default.jpg" {
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
				"error":  "can not register",
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
	}
	token, err := generateToken(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"msg":    "can not generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "register successfully",
		"token":   token,
	})
}
