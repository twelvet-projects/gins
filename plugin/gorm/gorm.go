package gorm

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/plugin/gorm/config"
	"github.com/twelvet-s/gins/plugin/gorm/initialize"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config // 配置

	INSTANCE_DB *gorm.DB // GORM实例

	INSTANCE_DBList map[string]*gorm.DB // 多数据源实例
)

type PluginGorm struct{}

func CreateGormPlug(config *config.Config) *PluginGorm {
	CONFIG = config
	return &PluginGorm{}
}

func (*PluginGorm) Register(group *gin.RouterGroup) {
	InitGorm()
}

func (*PluginGorm) RouterPath() string {
	return "gorm"
}

// InitGorm 初始化数据库并产生数据库全局变量
func InitGorm() {
	dbConfig := CONFIG.Datasource
	switch CONFIG.DbType {
	case "mysql":
		INSTANCE_DB = initialize.GormMysql(config.Mysql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "pgsql":
		INSTANCE_DB = initialize.GormPgSql(config.Pgsql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "oracle":
		INSTANCE_DB = initialize.GormOracle(config.Oracle{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "mssql":
		INSTANCE_DB = initialize.GormMssql(config.Mssql{
			GeneralDB: dbConfig.GeneralDB,
		})
	case "sqlite":
		INSTANCE_DB = initialize.GormSqlite(config.Sqlite{
			GeneralDB: dbConfig.GeneralDB,
		})
	default:
		INSTANCE_DB = initialize.GormMysql(config.Mysql{
			GeneralDB: dbConfig.GeneralDB,
		})
	}

	// 多数据源
	DBList()
}
