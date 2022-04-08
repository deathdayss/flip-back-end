package service

import (
	"container/list"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSecurityQuestion(c *gin.Context) {

	questionlist := list.New()
	questionlist.PushBack("sex")
	questionlist.PushBack("birth")
	questionlist.PushBack("phone")

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"question": questionlist,
	})
}
