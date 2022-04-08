package service

import (
	"container/list"
	"net/http"
	"os"
	"strings"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary register a new account
// @Description using password, email and nickname to create a new account
// @Accept  plain
// @Produce  json
// @Param   email     body    string     true        "email"
// @Param   password     body    string     true        "password"
// @Param   nickname     body    string     true        "nickname"
// @Param   file_body     body    string     false        "person image"
// @Success 200 {string} json   "{"status":200, "msg":"register successfully":, token":"string"}"
// @Failure 406 {string} json	"email has been used"
// @Failure 406 {string} json	"email, nickname or password is missing"
// @Failure 400 {string} json	"cannot save answer"
// @Failure 401 {string} json	"can not generate token"
// @Router /v1/notoken/register [post]
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

	questionlist1 := list.New()
	questionlist1.PushBack("sex")
	questionlist1.PushBack("birth")
	questionlist1.PushBack("phone")

	questionlist2 := list.New()
	questionlist2.PushBack("sex")
	questionlist2.PushBack("birth")
	questionlist2.PushBack("phone")

	questionlist3 := list.New()
	questionlist3.PushBack("sex")
	questionlist3.PushBack("birth")
	questionlist3.PushBack("phone")

	c.JSON(http.StatusOK, gin.H{
		"question1": questionlist1,
		"question2": questionlist2,
		"question3": questionlist3,
	})

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
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"msg":    "can not generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "register successfully",
		"token": token,
	})
}
