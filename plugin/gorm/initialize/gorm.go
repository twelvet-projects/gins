package initialize

import (
	"database/sql"
	"github.com/twelvet-s/gins/application/system/model"
	"github.com/twelvet-s/gins/framework/global"
	"github.com/twelvet-s/gins/plugin/gorm"
	"github.com/twelvet-s/gins/plugin/gorm/config"
	globalPugin "github.com/twelvet-s/gins/plugin/gorm/global"
	"os"

	"go.uber.org/zap"
)

// GormInit 初始化数据库并产生数据库全局变量
func GormInit() {
	dbConfig := globalPugin.CONFIG.Datasource
	switch globalPugin.CONFIG.DbType {
	case "mysql":
		gorm.INSTANCE_DB = GormMysql(config.Mysql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "pgsql":
		gorm.INSTANCE_DB = GormPgSql(config.Pgsql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "oracle":
		gorm.INSTANCE_DB = GormOracle(config.Oracle{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "mssql":
		gorm.INSTANCE_DB = GormMssql(config.Mssql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "sqlite":
		gorm.INSTANCE_DB = GormSqlite(config.Sqlite{
			GeneralDB: dbConfig.GeneralDB,
		})
	default:
		gorm.INSTANCE_DB = GormMysql(config.Mysql{
			GeneralDB: dbConfig.GeneralDB,
		})
	}

	// 多数据源
	DBList()

	// 程序结束前关闭数据库链接
	if gorm.INSTANCE_DB != nil {
		RegisterTables() // 初始化表
		db, _ := gorm.INSTANCE_DB.DB()
		defer func(db *sql.DB) {
			_ = db.Close()
		}(db)
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
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
}
