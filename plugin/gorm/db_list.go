package gorm

import (
	"github.com/twelvet-s/gins/plugin/gorm/config"
	"github.com/twelvet-s/gins/plugin/gorm/initialize"
	"gorm.io/gorm"
)

const sys = "system"

// DBList 多数据源加载
func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = initialize.GormMysql(config.Mysql{GeneralDB: info.GeneralDB})
		case "mssql":
			dbMap[info.AliasName] = initialize.GormMssql(config.Mssql{GeneralDB: info.GeneralDB})
		case "pgsql":
			dbMap[info.AliasName] = initialize.GormPgSql(config.Pgsql{GeneralDB: info.GeneralDB})
		case "oracle":
			dbMap[info.AliasName] = initialize.GormOracle(config.Oracle{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	// 做特殊判断,是否有迁移
	// 适配低版本迁移多数据库版本
	if sysDB, ok := dbMap[sys]; ok {
		INSTANCE_DB = sysDB
	}
	INSTANCE_DBList = dbMap
}
