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
		noTokenFeature.GET("/get/product", service.GetProductInfo)
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
		SearchFeature.GET("/item/:mode", middleware.Auth(), service.Search)
		SearchFeature.GET("/rank/:mode", middleware.Auth(), service.SearchRank)
		SearchFeature.GET("/notoken/item/:mode", service.Search)
		SearchFeature.GET("/notoken/rank/:mode", service.SearchRank)
		SearchFeature.GET("/history", middleware.Auth(), service.GetSearchHistory)
	}
	PersonalZoneFeature := engine.Group("/v1/personalzone")
	{
		PersonalZoneFeature.Use(middleware.Auth())
		PersonalZoneFeature.GET("/product", service.GetPersonalProduct)
		PersonalZoneFeature.POST("/changeico", service.ChangeIcon)
		PersonalZoneFeature.POST("/replace", service.ChangeFile)
	}

	MultiZoneFeature := engine.Group("/v2")
	{

		//search
		MultiZoneFeature.GET("/search/game", service.SearchGameByZone)
	}

	MultiZoneFeature_UpLoad := engine.Group("/v2/upload")
	{
		MultiZoneFeature_UpLoad.Use(middleware.Auth())
		MultiZoneFeature_UpLoad.POST("/info", service.UploadInfoByZone)
		MultiZoneFeature_UpLoad.POST("/game", service.UploadZipByZone)
	}

	MultiZoneFeature_Download := engine.Group("/v2/download")
	{
		MultiZoneFeature_Download.GET("/img", service.DownloadImgByZone)
		MultiZoneFeature_Download.GET("/game", service.DownloadGameByZone)
		MultiZoneFeature_Download.GET("/personal", service.DownloadPersonalByZone)
	}

	MultiZoneFeature_Ranking := engine.Group("/v2/rank")
	{
		//getRanking
		MultiZoneFeature_Ranking.GET("/zone", service.GetGameRankingByMultiZone)
		MultiZoneFeature_Ranking.GET("/download", service.GetGameRankingDownloadingByZone)
		MultiZoneFeature_Ranking.GET("/author", service.GetAurthorRankingByZone)
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(":8084")
	return engine
}
