package service

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	email, ok1 := c.GetQuery("email")
	password, ok2 := c.GetQuery("password")
	if !ok1 || !ok2 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email or password is missing",
		})
		return
	}
	if repository.CheckExistence(email) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email is used",
		})
		return
	}
	if !checkPassword(password) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "the password does not comply the rule, please set again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "the register is legal",
	})
}

func checkPassword(password string) bool {
	cntCapital, cntLower, cntNum := 0, 0, 0
	for i := range password {
		if password[i] >= 'a' && password[i] <= 'z' {
			cntLower++
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			cntCapital++
		} else if password[i] >= '0' && password[i] <= '9' {
			cntNum++
		}
	}
	return cntCapital > 0 && cntLower > 0 && cntNum > 0
}

// @Summary register a new account
// @Description using password, email and nickname to create a new account
// @Accept  plain
// @Produce  json
// @Param   email     body    string     true        "email"
// @Param   password     body    string     true        "password"
// @Param   nickname     body    string     true        "nickname"
// @Param   question1     body    int    true        "question1"
// @Param   answer1    body    string     true        "answer1"
// @Param   question2     body    int    true        "question2"
// @Param   answer2    body    string     true        "answer2"
// @Param   question3     body    int    true        "question3"
// @Param   answer3    body    string     true        "answer3"
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

	question1, ok4 := c.GetPostForm("question1")
	answer1, ok5 := c.GetPostForm("answer1")
	question2, ok6 := c.GetPostForm("question2")
	answer2, ok7 := c.GetPostForm("answer2")
	question3, ok8 := c.GetPostForm("question3")
	answer3, ok9 := c.GetPostForm("answer3")
	question1num, err1 := strconv.Atoi(question1)
	question2num, err2 := strconv.Atoi(question2)
	question3num, err3 := strconv.Atoi(question3)

	if !ok4 || !ok5 || !ok6 || !ok7 || !ok8 || !ok9 || len(answer1) == 0 || len(answer2) == 0 || len(answer3) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "answer or question missing",
		})
		return
	}

	if question1num == question2num || question1num == question3num || question2num == question3num {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "do not select repeated questions",
		})
		return
	}

	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "Invalid id",
		})
		return
	}

	if !repository.CheckAnswer(email) {
		_, err := repository.AddAnswer(email, answer1, answer2, answer3, question1num, question2num, question3num)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "cannot save answer",
			})
			return
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
		"token":   token,
	})
}
