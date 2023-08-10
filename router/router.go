package router

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/application/system/model"
	system "github.com/twelvet-s/gins/application/system/router"
	"github.com/twelvet-s/gins/framework/global"
	"github.com/twelvet-s/gins/framework/initialize"
	"github.com/twelvet-s/gins/plugin/gorm"
	"go.uber.org/zap"
	"os"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	router := gin.Default()

	// 安装插件
	initialize.InstallPlugin(router)

	db := gorm.INSTANCE_DB
	err := db.AutoMigrate(
		// 系统模块表
		model.SysDept{},
		model.SysDfs{},
		model.SysDictData{},
		model.SysDictType{},
		model.SysJob{},
		model.SysJobLog{},
		model.SysLoginInfo{},
		model.SysMenu{},
		model.SysOperationLog{},
		model.SysPost{},
		model.SysRole{},
		model.SysRoleDept{},
		model.SysRoleMenu{},
		model.SysUser{},
		model.SysUserPost{},
		model.SysUserRole{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")

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
