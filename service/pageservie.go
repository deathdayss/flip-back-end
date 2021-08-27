package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pageservice(c *gin.Context) {

	pages = page.repository.Rank()
	if len(pages) != 0 {
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "cannot find relevant recommandation",
	})

}
