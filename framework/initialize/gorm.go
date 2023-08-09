package initialize

import (
	"database/sql"
	"github.com/twelvet-s/gins/application/system/model"
	"github.com/twelvet-s/gins/config"
	"github.com/twelvet-s/gins/framework/global"
	"os"

	"go.uber.org/zap"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() {
	dbConfig := global.CONFIG.Gins.Datasource
	switch global.CONFIG.Gins.DbType {
	case "mysql":
		global.DB = GormMysql(config.Mysql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "pgsql":
		global.DB = GormPgSql(config.Pgsql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "oracle":
		global.DB = GormOracle(config.Oracle{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "mssql":
		global.DB = GormMssql(config.Mssql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "sqlite":
		global.DB = GormSqlite(config.Sqlite{
			GeneralDB: dbConfig.GeneralDB,
		})
	default:
		global.DB = GormMysql(config.Mysql{
			GeneralDB: dbConfig.GeneralDB,
		})
	}

	// 多数据源
	DBList()

	// 程序结束前关闭数据库链接
	if global.DB != nil {
		RegisterTables() // 初始化表
		db, _ := global.DB.DB()
		defer func(db *sql.DB) {
			_ = db.Close()
		}(db)
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.DB
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
