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

// @Summary get security code
// @Description get a random security code
// @Accept  plain
// @Produce  json
// @Param   getCode     header    int     true        "getCode"
// @Success 200 {string} json  "{"status":200, "content":code.Content,"url":code,URL}"
// @Router /v1/verification/code [GET]
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
