package main

import (
	"github.com/twelvet-s/gins/framework"
	"github.com/twelvet-s/gins/framework/global"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gins API
// @version                     1.0.0
// @description                 Gins Swagger
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 初始化Viper
	global.Viper = framework.Viper()
	// 初始化zap日志库
	global.LOG = framework.Zap()
	zap.ReplaceGlobals(global.LOG)
	// 启动服务
	framework.StartServer()
}
