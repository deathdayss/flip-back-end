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
		rankFeature.GET("/download", service.GetGameRankingDownloading)
	}
	userinfoFeature := engine.Group("/v1/info")
	{
		userinfoFeature.POST("/getuserinfo", service.GetUserInfo)
	}
	downloadFeature := engine.Group("/v1/download")
	{
		downloadFeature.GET("/img", service.DownloadImg)
	}
	uploadFeature := engine.Group("/v1/upload")
	{
		uploadFeature.POST("/img", service.UploadImg)
		uploadFeature.POST("/game", service.UploadZip)
	}
	likeFeature := engine.Group("/v1/like")
	{
		likeFeature.GET("/click", service.LikeOrUnlike)
		likeFeature.GET("/check", service.HasLike)
		likeFeature.GET("/num", service.GetLikeNum)
	}

	shareFeature := engine.Group("/v1/share")
	{
		shareFeature.GET("/share", service.ShareOperation)
	}

	getInfoFeature := engine.Group("/v1/get")
	{
		getInfoFeature.GET("/product", service.GetProductInfo)
	}

	verificationCodeFeature := engine.Group("/v1/verification")
	{
		verificationCodeFeature.GET("/code", service.GetCode)
	}

	return engine
}
