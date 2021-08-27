package router

import (
	"D/flip-back-end/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageRouter(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The router is wrong.")
	})
	userFeature := engine.Group("/v1/user")
	{
		userFeature.POST("/page", service.Pageservice) // /v1/user/register
	}
	return engine
}
