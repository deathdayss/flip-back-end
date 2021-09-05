package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func DownloadImg(c *gin.Context) {
	filename := c.Query("img_name")
	ss := strings.Split(filename, ".")
	if len(ss) < 2 || (strings.ToLower(ss[1]) != "jpg" && strings.ToLower(ss[1]) != "jpeg" && strings.ToLower(ss[1]) != "png") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	contentType := ss[1]
	if _, err := os.Stat("./storage/thumbnail/"+filename); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": "the img name is wrong",
		})
		return
	}
	if contentType == "jpg" || contentType == "jpeg" {
		c.Writer.Header().Add("Content-Type", "image/jpeg")
	} else {
		c.Writer.Header().Add("Content-Type", "image/png")
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("inline;filename=%s", filename))
	c.File("./storage/thumbnail/"+filename)
}