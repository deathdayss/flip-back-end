package main

import (
	"github.com/deathdayss/flip-back-end/models"
	"github.com/deathdayss/flip-back-end/router"
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

