package main

import (
	"D/flip-back-end/models"
	"D/flip-back-end/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.DbClient.Init()
	defer models.DbClient.Close()
	r := gin.New()
	router.RegisterRouter(r)
	r.Run(":8084")
}

