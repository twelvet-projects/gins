package router

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/router/index"
)

func Init() *gin.Engine {

	router := gin.Default()

	// 处理500
	router.Use(Advice500)

	// 处理404
	router.NoRoute(Advice404)

	// 首页
	homeRouter := index.HomeRouter{}
	systemPrivateGroup := router.Group("/")
	{
		// 首页
		homeRouter.InitRouter(systemPrivateGroup)
	}

	return router
}
