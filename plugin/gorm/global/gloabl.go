package global

import (
	"github.com/twelvet-s/gins/plugin/gorm/config"
	"gorm.io/gorm"
)

var (
	CONFIG config.Config // 配置

	DBList map[string]*gorm.DB // 多数据源

	DB *gorm.DB // 数据库
)
