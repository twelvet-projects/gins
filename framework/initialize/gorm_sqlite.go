package initialize

import (
	"github.com/glebarez/sqlite"
	"github.com/twelvet-s/gins/config"
	"github.com/twelvet-s/gins/framework/initialize/internal"
	"gorm.io/gorm"
)

// GormSqlite 初始化Sqlite数据库
func GormSqlite(s config.Sqlite) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}
