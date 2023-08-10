package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/framework/global"
	"github.com/twelvet-s/gins/framework/utils/plugin"
	"github.com/twelvet-s/gins/plugin/gorm"
	"github.com/twelvet-s/gins/plugin/redis"
)

// PluginInit 进行初始化
func PluginInit(group *gin.RouterGroup, plugins ...plugin.Plugin) {
	for _, pluginItem := range plugins {
		pluginGroup := group.Group(pluginItem.RouterPath())
		pluginItem.Register(pluginGroup)
	}
}

// InstallPlugin 安装插件
func InstallPlugin(router *gin.Engine) {
	publicGroup := router.Group("")
	fmt.Println("无鉴权插件安装==》", publicGroup)
	privateGroup := router.Group("")

	fmt.Println("鉴权插件安装==》", privateGroup)

	// 加载Gorm
	PluginInit(privateGroup, gorm.CreateGormPlug(&global.CONFIG.Gins.Gorm))
	// 加载Redis
	PluginInit(privateGroup, redis.CreateRedisPlug(&global.CONFIG.Gins.Redis))
}
