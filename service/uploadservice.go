package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

func(c *gin.Context){

	form, _ := c.MultipartForm()

	files := form.File["upload[]"]

	for _, file := range files {
		c.SaveUploadedFile(file, file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d upload success!", len(files)))
}
