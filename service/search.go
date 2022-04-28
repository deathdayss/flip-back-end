package service

import (
	"net/http"
	"strconv"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

var AllowedGameMtd map[string]bool = map[string]bool{"like": true, "download": true, "comment": true, "time": true}

// @Summary search a game by keyword
// @Description search a game by keyword
// @Accept  plain
// @Produce  json
// @Param   num     header    int     true        "the number of the return item"
// @Param   keyword     header    string     true        "the keyword"
// @Param   method  header     string true "the order method"
// @Param   offset  header     int true "the offset"
// @Success 200 {array} dto.RankItem   "{"status":200, "List":list}"
// @Router /v1/search/game [GET]
func SearchPerson(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	keyword := c.Query("keyword")
	zone := c.Query("zone")
	if err != nil || len(keyword) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "keyword or num is missing",
		})
		return
	}
	mode := c.Query("mode")
	if mode == "" {
		mode = "game"
	}
	if err := repository.AddWord(keyword, mode); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "can not connect the redis db",
		})
		return
	}
	rankMtd, ok := c.GetQuery("method")
	if !ok || (rankMtd != "like" && rankMtd != "download" && rankMtd != "comment") {
		rankMtd = "like"
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
		offset = offset * (num - 1)
	}
	rankInfo, err := repository.SearchGame(keyword, num, offset, rankMtd, zone)
	if err != nil || len(*rankInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}
	rankList := []dto.RankItem{}
	for _, ri := range *rankInfo {
		/*
			fp, err := ioutil.ReadFile("./storage/thumbnail/"+ri.ImgUrl)
			if err != nil {
				fp, _ = ioutil.ReadFile("./storage/thumbnail/not_found.png")
			}*/
		rankList = append(rankList, dto.RankItem{
			ID:          ri.ID,
			Name:        ri.Name,
			LikeNum:     ri.LikeNum,
			DownloadNum: ri.DownloadNum,
			CommentNum:  ri.CommentNum,
			Img:         ri.ImgUrl,
			AuthorName:  repository.FindNickName(ri.UID),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List":   rankList,
	})
}

// @Summary search a game by keyword
// @Description search a game by keyword
// @Accept  plain
// @Produce  json
// @Param   num     header    int     true        "the number of the return item"
// @Param   keyword     header    string     true        "the keyword"
// @Param   method  header     string true "the order method"
// @Param   offset  header     int true "the offset"
// @Success 200 {array} dto.RankItem   "{"status":200, "List":list}"
// @Router /v1/search/:mode [GET]
func Search(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	keyword := c.Query("keyword")
	zone := c.Query("zone")
	if err != nil || len(keyword) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "keyword or num is missing",
		})
		return
	}
	email, ok := c.Get("email")
	if ok {
		err := repository.AddSearchHistory(email.(string), keyword)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 406,
				"error": "redis server can not add the search key to history",
			})
			return
		}
	}
	mode := c.Param("mode")
	if mode != "game" && mode != "person" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "can not search this sort",
		})
		return
	}
	if err := repository.AddWord(keyword, mode); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "can not connect the redis db",
		})
		return
	}
	rankMtd, ok := c.GetQuery("method")
	if !ok || (rankMtd != "like" && rankMtd != "download" && rankMtd != "comment") {
		rankMtd = "like"
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
		offset = offset * (num - 1)
	}
	if mode == "person" {
		rankInfo, err := repository.SearchPerson(keyword, num, offset, rankMtd)
		if err != nil || len(*rankInfo) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 404,
				"error":  "no data",
			})
			return
		}
		rankList := make([]int, len(*rankInfo))
		for i := range rankList {
			rankList[i] = (*rankInfo)[i].ID
		}
		c.JSON(http.StatusOK, gin.H{
			"Status": 200,
			"List":   rankList,
		})
		return
	}
	rankInfo, err := repository.SearchGame(keyword, num, offset, rankMtd, zone)
	if err != nil || len(*rankInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}
	rankList := []dto.RankItem{}
	for _, ri := range *rankInfo {
		/*
			fp, err := ioutil.ReadFile("./storage/thumbnail/"+ri.ImgUrl)
			if err != nil {
				fp, _ = ioutil.ReadFile("./storage/thumbnail/not_found.png")
			}*/
		rankList = append(rankList, dto.RankItem{
			ID:          ri.ID,
			Name:        ri.Name,
			LikeNum:     ri.LikeNum,
			DownloadNum: ri.DownloadNum,
			CommentNum:  ri.CommentNum,
			Img:         ri.ImgUrl,
			AuthorName:  repository.FindNickName(ri.UID),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List":   rankList,
	})
}

func GetSearchHistory(c *gin.Context) {
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "email is missing",
		})
		return
	}
	history, err := repository.GetSearchHistory(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 406,
			"error":  "redis server error ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"history": history,
	})
}

func SearchRank(c *gin.Context) {
	mode := c.Param("mode")
	if mode != "person" && mode != "game" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "can not search this sort",
		})
		return
	}
	searchAns, err := repository.SearchRank(mode)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "internal error in redis db",
		})
		return
	}
	if len(searchAns) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"words":  searchAns,
	})
}
