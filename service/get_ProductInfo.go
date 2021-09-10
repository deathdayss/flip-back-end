package service

import (
	"net/http"
	//"io/ioutil"
	"strconv"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"

	//"github.com/deathdayss/flip-back-end/dto"
	"github.com/gin-gonic/gin"
)

func GetProductInfo(c *gin.Context) { //gin.Context用于处理http请求
	pid, err := strconv.Atoi(c.Query("pid")) //获取请求的pid，把String转为int
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "The pid is illegal",
		})
		return
	}

	if !repository.CheckGameExistence(pid) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "No such game",
		})
		return
	}

	productInfo := dto.ProductItem{}

	return productInfo

}
