package initialize

import (
	"github.com/twelvet-s/gins/global"
	"github.com/twelvet-s/gins/model/system"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.CONFIG.Gins.DbType {
	case "mysql":
		return GormMysql()
	///*case "pgsql":
	//	return GormPgSql()
	//case "oracle":
	//	return GormOracle()
	//case "mssql":
	//	return GormMssql()
	//case "sqlite":
	//	return GormSqlite()*/
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		// 系统模块表
		system.SysUser{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
