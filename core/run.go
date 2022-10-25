package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/g"
	"github.com/twelvet-s/gins/router"
	"go.uber.org/zap"
)

func Run() {
	gin.SetMode(gin.DebugMode)
	g.VP = Viper() // 初始化Viper
	g.LOG = Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(g.LOG)

	// 程序结束前关闭数据库链接
	if g.DB != nil {
		db, _ := g.DB.DB()
		defer db.Close()
	}

	// 初始化路由
	routerInit := router.Init()

	// 启动服务
	port := fmt.Sprintf(":%d", g.CONFIG.Server.Port)
	server := StartServer(port, routerInit)
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
