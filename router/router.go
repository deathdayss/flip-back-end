package router

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/service"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The router is wrong.")
	})
	userFeature := engine.Group("/v1/user") 
	{
		userFeature.POST("/register", service.Register) // /v1/user/register
		userFeature.POST("/login", service.Login)
	}
	rankFeature := engine.Group("/v1/rank")
	{
		rankFeature.GET("/zone", service.GetGameRanking)
	}
	return engine
}