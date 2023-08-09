package initialize

import (
	//"github.com/dzwvip/oracle"
	"github.com/twelvet-s/gins/config"
	"github.com/twelvet-s/gins/framework/initialize/internal"
	//_ "github.com/godror/godror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormOracle 初始化oracle数据库
// 需要 GCC 支持，否则无法运行
// 如果需要Oracle库 放开import里的注释 把下方 mysql.Config 改为 oracle.Config ;  mysql.New 改为 oracle.New
func GormOracle(o config.Oracle) *gorm.DB {
	if o.Dbname == "" {
		return nil
	}
	oracleConfig := mysql.Config{
		DSN:               o.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(mysql.New(oracleConfig), internal.Gorm.Config(o.Prefix, o.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(o.MaxIdleConns)
		sqlDB.SetMaxOpenConns(o.MaxOpenConns)
		return db
	}
}
