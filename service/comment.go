package service

import (
	"net/http"
	"strconv"

	"github.com/deathdayss/flip-back-end/dto"
	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

// @Summary update a comment
// @Description update a comment
// @Accept  plain
// @Produce  json
// @Param   email    body    string     true        "email"
// @Param   comment_id     body    int     true        "comment_id"
// @Success 200 {string} json   "{"status":200, "message":"the comment has been up"}"
// @Router /v1/change/commentcomment/up [POST]
func UpComment(c *gin.Context) {
	emailIt, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "unauth token",
		})
		return
	}
	email := emailIt.(string)
	pid, err := repository.Email2PID(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "can not auth the email",
		})
		return
	}
	cidstr, ok := c.GetPostForm("comment_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "the comment id does not exist",
		})
		return
	}
	cid, err := strconv.Atoi(cidstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "the comment id does not exist",
		})
		return
	}
	err = repository.UpComment(cid, pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "the comment has been up by the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "the comment has been up",
	})
}

// @Summary search a game by keyword
// @Description search a game by keyword
// @Accept  plain
// @Produce  json
// @Param   num     header    int     true        "the number of the return item"
// @Param   GID     header    string     true        "GID"
// @Param   offset  header     int true "the offset"
// @Param   method  header     string true "the order method"
// @Success 200 {array} dto.RankComment  "{"status":200, "List":ranklist}"
// @Router /v1/rank/comment/time [GET]
func GetCommentRanking(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("num"))
	zone := c.Query("GID")
	if err != nil || len(zone) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": 406,
			"error":  "num or zone is wrong",
		})
		return
	}
	rankMtd, ok := c.GetQuery("method")
	if !ok || (rankMtd != "time" && rankMtd != "like") {
		rankMtd = "time"
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
		offset = offset * (num-1)
	}
	rankInfo, err := repository.GetCommentRanking(zone, num, rankMtd, offset)
	if err != nil || len(*rankInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "no data",
		})
		return
	}
	rankList := []dto.RankComment{}
	for _, ri := range *rankInfo {
		/*
			fp, err := ioutil.ReadFile("./storage/thumbnail/"+ri.ImgUrl)
			if err != nil {
				fp, _ = ioutil.ReadFile("./storage/thumbnail/not_found.png")
			}*/
		rankList = append(rankList, dto.RankComment{
			ID:         ri.ID,
			Comment:    ri.Content,
			LikeNum:    ri.Up,
			CreateTime: ri.CreateTime.String(),
			GID:        ri.GID,
			UID:        ri.UID,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": 200,
		"List":   rankList,
	})
}

// @Summary add a comment
// @Description add a comment
// @Accept  plain
// @Produce  json
// @Param   email    body    string     true        "email"
// @Param   game_id     body    int     true        "gameid"
// @Success 200 {string} json   "{"status":200, "message":the comment has bee added}"
// @Router /v1/change/comment/add [POST]
func AddComment(c *gin.Context) {
	emailIt, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "unauth token",
		})
		return
	}
	email := emailIt.(string)
	pid, err := repository.Email2PID(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "can not auth the email",
		})
		return
	}
	gidstr, ok := c.GetPostForm("game_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "the game id does not exist",
		})
		return
	}
	gid, err := strconv.Atoi(gidstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "the game id does not exist",
		})
		return
	}
	comment, ok := c.GetPostForm("comment")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "the game id does not exist",
		})
		return
	}
	err = repository.AddComment(comment, gid, pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "the comment can not be added",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "the comment has been added",
	})
}
