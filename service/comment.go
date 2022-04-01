package service

import (
	"net/http"
	"strconv"

	"github.com/deathdayss/flip-back-end/repository"
	"github.com/gin-gonic/gin"
)

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
			"error": "the game id does not exist",
		})
		return
	}
	gid, err := strconv.Atoi(gidstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error": "the game id does not exist",
		})
		return 
	}
	comment, ok := c.GetPostForm("comment")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error": "the game id does not exist",
		})
		return 
	}
	err = repository.AddComment(comment, gid, pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error": "the comment can not be added",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg": "the comment has been added",
	})
}