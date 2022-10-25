package router

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/router/system"
)

type GinsRouterGroup struct {
	System system.RouterGroup
}

func Init() *gin.Engine {

	router := gin.Default()

	// 处理500
	router.Use(Advice500)

	// 处理404
	router.NoRoute(Advice404)

	routerGroup := new(GinsRouterGroup)

	// 设置静态目录
	router.Static("/static", "./resource/static")
	// favicon.ico
	router.StaticFile("/favicon.ico", "./resource/static/favicon.ico")

	// 系统API服务组
	systemGroup := routerGroup.System
	systemPrivateGroup := router.Group("/system")
	{
		// 用户路由
		systemGroup.InitUserRouter(systemPrivateGroup)
		// 角色路由
		systemGroup.InitRoleRouter(systemPrivateGroup)
	}

	return router
}
