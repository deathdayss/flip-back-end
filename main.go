package main

import (
	"github.com/deathdayss/flip-back-end/models"
	"github.com/deathdayss/flip-back-end/router"
	"github.com/deathdayss/flip-back-end/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.DbClient.Init()
	r := gin.New()
	router.RegisterRouter(r, 
		middleware.Cors(),
	)
	r.Run(":8084")
}

