package framework

import (
	"fmt"
	"github.com/twelvet-s/gins/framework/global"
	"github.com/twelvet-s/gins/framework/initialize"
	"github.com/twelvet-s/gins/router"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
}

// StartServer 启动服务
func StartServer() {

	// 初始化路由
	routerInit := router.InitRouter()

	// 初始化数据
	initialize.InitDataSource()

	// 启动服务
	port := fmt.Sprintf(":%d", global.CONFIG.Server.Port)

	server := http.Server{
		Addr:           port,
		Handler:        routerInit,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

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
