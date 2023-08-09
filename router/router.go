package router

import (
	"github.com/gin-gonic/gin"
	system "github.com/twelvet-s/gins/application/system/router"
	"github.com/twelvet-s/gins/framework/global"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	router := gin.Default()

	// 处理500
	router.Use(Advice500)

	// 处理404
	router.NoRoute(Advice404)

	// 系统模块
	systemRouter := system.SystemRouterGroupApp

	// 首页
	systemPrivateGroup := router.Group(global.CONFIG.Server.RouterPrefix)
	{
		// 首页
		systemRouter.InitIndexRouter(systemPrivateGroup)
	}

	return router
}
