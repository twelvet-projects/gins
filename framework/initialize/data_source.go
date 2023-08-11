package initialize

import (
	"github.com/twelvet-s/gins/application/system/model"
	"github.com/twelvet-s/gins/framework/global"
	"github.com/twelvet-s/gins/plugin/gorm"
	"go.uber.org/zap"
	"os"
)

// InitDataSource 初始化表
func InitDataSource() {
	db := gorm.INSTANCE_DB
	if db != nil {
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
	}

}
