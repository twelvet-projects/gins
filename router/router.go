package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	system "github.com/twelvet-s/gins/application/system/router"
	"github.com/twelvet-s/gins/docs"
	"github.com/twelvet-s/gins/framework/global"
	"github.com/twelvet-s/gins/framework/initialize"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	router := gin.Default()

	// 安装插件
	initialize.InstallPlugin(router)

	// 处理500
	router.Use(Advice500)

	// 处理404
	router.NoRoute(Advice404)

	// Swagger
	docs.SwaggerInfo.BasePath = global.CONFIG.Server.RouterPrefix
	router.GET(fmt.Sprintf("%s/swagger/*any", global.CONFIG.Server.RouterPrefix), ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置静态目录
	router.Static(fmt.Sprintf("%s/static", global.CONFIG.Server.RouterPrefix), "./resources/static")
	// favicon.ico
	router.StaticFile(fmt.Sprintf("%s/favicon.ico", global.CONFIG.Server.RouterPrefix), "./resources/static/favicon.ico")

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
