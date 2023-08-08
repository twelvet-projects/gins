package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/framework"
	"github.com/twelvet-s/gins/global"
	"github.com/twelvet-s/gins/initialize"
	"github.com/twelvet-s/gins/router"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	gin.SetMode(gin.DebugMode)
	// 初始化Viper
	global.Viper = framework.Viper()
	// 初始化zap日志库
	global.LOG = framework.Zap()
	zap.ReplaceGlobals(global.LOG)

	// gorm连接数据库
	initialize.Gorm()
	// redis
	initialize.Redis()

	// 初始化路由
	routerInit := router.Init()
	// 设置静态目录
	routerInit.Static("/static", "./resources/static")
	// favicon.ico
	routerInit.StaticFile("/favicon.ico", "./resources/static/favicon.ico")

	// 启动服务
	port := fmt.Sprintf(":%d", global.CONFIG.Server.Port)
	server := framework.StartServer(port, routerInit)
	fmt.Printf(`
	   _______           
	  / ____(_)___  _____
	 / / __/ / __ \/ ___/
	/ /_/ / / / / (__  ) 
	\____/_/_/ /_/____/

	欢迎使用 Gins
	当前版本: v1.0.0
	加群方式: QQ群：985830229
	自动化文档地址: http://127.0.0.1%s/swagger/index.html
	API地址: http://127.0.0.1%s

`,
		port, port)
	server.ListenAndServe().Error()
}
