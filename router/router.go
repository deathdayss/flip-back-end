package router

import (
	"net/http"

	"github.com/deathdayss/flip-back-end/middleware"
	"github.com/deathdayss/flip-back-end/service"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"
)

// swagger embed files

// @title           FLIP backend API
// @version         1.0
// @description     FLIP backend server.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8084

// @securityDefinitions.basic  JWT
func RegisterRouter(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The router is wrong.")
	})
	noTokenFeature := engine.Group("/v1/notoken")
	{
		noTokenFeature.POST("/register", service.Register) // /v1/user/register
		noTokenFeature.POST("/login", service.Login)
		noTokenFeature.GET("/verify", service.Verify)
		noTokenFeature.POST("/change/vertify", service.VertifyExist)
		noTokenFeature.POST("/change/answer", service.VertifyAnswer)
		noTokenFeature.POST("/change/password", service.ChangePassword)
	}
	userFeature := engine.Group("/v1/user")
	{
		//userFeature.POST("/register", service.Register) // /v1/user/register
		//userFeature.POST("/login", service.Login)
		userFeature.Use(middleware.Auth())
		userFeature.GET("/detail", service.GetUserDetail)
		userFeature.POST("/change/detail", service.ChangeDetail)

	}
	dataFeature := engine.Group("/v1/data")
	dataFeature.Use(middleware.Auth())
	{
		//dataFeature.Use(middleware.Auth())
		dataFeature.GET("/datebytime", service.GetDataByTime)
	}
	rankFeature := engine.Group("/v1/rank")
	{
		//rankFeature.Use(middleware.Auth())
		rankFeature.GET("/zone", service.GetGameRanking)
		rankFeature.GET("/download", service.GetGameRankingDownloading)
		rankFeature.GET("/author", service.GetAurthorRankingByZone)
		//newRank Settings
		rankFeature.GET("/multi_zone", service.GetGameRankingByMultiZone)
	}
	userinfoFeature := engine.Group("/v1/info")
	{
		userFeature.Use(middleware.Auth())
		userinfoFeature.POST("/getuserinfo", service.GetUserInfo)
	}
	downloadFeature := engine.Group("/v1/download")
	{
		//downloadFeature.Use(middleware.Auth())
		downloadFeature.GET("/img", service.DownloadImg)
		downloadFeature.GET("/game", service.DownloadGame)
		downloadFeature.GET("/personal", service.DownloadPersonal)
	}
	uploadFeature := engine.Group("/v1/upload")
	{
		uploadFeature.Use(middleware.Auth())
		uploadFeature.POST("/info", service.UploadInfo)
		uploadFeature.POST("/game", service.UploadZip)
	}
	likeFeature := engine.Group("/v1/like")
	{
		likeFeature.Use(middleware.Auth())
		likeFeature.GET("/click", service.LikeOrUnlike)
		likeFeature.GET("/check", service.HasLike)
		likeFeature.GET("/num", service.GetLikeNum)
	}
	collectFeature := engine.Group("/v1/collect")
	{
		collectFeature.Use(middleware.Auth())
		collectFeature.GET("/click", service.CollectOrUncollect)
		collectFeature.GET("/check", service.HasCollect)
		collectFeature.GET("/num", service.GetCollectNum)
	}

	verificationCodeFeature := engine.Group("/v1/verification")
	{
		verificationCodeFeature.GET("/code", service.GetCode)
	}
	ChangeCommentFeature := engine.Group("/v1/change/comment")
	{
		ChangeCommentFeature.Use(middleware.Auth())
		ChangeCommentFeature.POST("/add", service.AddComment)
		ChangeCommentFeature.POST("/up", service.UpComment)
	}
	ShowCommentFeature := engine.Group("/v1/rank/comment")
	{
		ShowCommentFeature.GET("/time", service.GetCommentRanking)
	}
	SecurityQuestionFeature := engine.Group("/v1/security")
	{
		SecurityQuestionFeature.GET("/question", service.GetSecurityQuestion)
		SecurityQuestionFeature.GET("/user/question", service.FindSecurityQuestion)
	}
	SearchFeature := engine.Group("/v1/search")
	{
		SearchFeature.GET("/game", service.SearchGame)
		SearchFeature.GET("/rank", service.SearchRank)
	}

	MultiZoneFeature := engine.Group("/v2")
	{
		//upload
		MultiZoneFeature.Use(middleware.Auth())
		MultiZoneFeature.POST("/upload/info", service.UploadInfoByZone)
		MultiZoneFeature.POST("/upload/game", service.UploadZipByZone)

		//download
		MultiZoneFeature.GET("/download/img", service.DownloadImgByZone)
		MultiZoneFeature.GET("/download/game", service.DownloadGameByZone)
		MultiZoneFeature.GET("/download/personal", service.DownloadPersonalByZone)

		//getRanking
		MultiZoneFeature.GET("/rank/zone", service.GetGameRankingByMultiZone)
		MultiZoneFeature.GET("/rank/download", service.GetGameRankingDownloadingByZone)

		//search
		MultiZoneFeature.GET("/search/game", service.SearchGameByZone)
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(":8084")
	return engine
}
