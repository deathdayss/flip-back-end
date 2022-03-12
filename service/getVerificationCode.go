package service

import (
	"math/rand"
	"net/http"
	"time"

	"strconv"
	//"io/ioutil"
	"github.com/deathdayss/flip-back-end/repository"

	//"github.com/deathdayss/flip-back-end/dto"
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	getCode, err := strconv.Atoi(c.Query("getCode"))
	if err != nil || getCode == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	ID := rand.Intn(10) + 1
	code, err := repository.GetCode(ID)

	c.JSON(http.StatusOK, gin.H{
		"Status":  200,
		"Content": code.Content,
		"URL":     code.URL,
	})

}
