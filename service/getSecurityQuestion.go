package service

import (
	"net/http"
	"strconv"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary get security question list
// @Description get security question list
// @Accept  plain
// @Produce  json
// @Param   num     header    int     true        "num"
// @Success 200 {object} dto.QuestionList  "{"status":200, "detail":questionlist}"
// @Router /v1/sequrity/question  [GET]

func GetSecurityQuestion(c *gin.Context) {

	num, err := strconv.Atoi(c.Query("num"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "num is wrong",
		})
		return
	}
	var offset int
	offsetStr, ok := c.GetQuery("offset")
	if !ok {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"status": 406,
				"error":  "offset if wrong",
			})
			return
		}
	}
	questionInfo, err := repository.GetQuestion(num, offset)
	if err != nil || len(*questionInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}

	questionList := []dto.QuestionItem{}
	for _, ri := range *questionInfo {
		questionList = append(questionList, dto.QuestionItem{
			ID:      ri.ID,
			Content: ri.Content,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List":   questionList,
	})
}
