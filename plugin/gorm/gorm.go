package gorm

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/plugin/gorm/initialize"
	"gorm.io/gorm"
)

var (
	GORM *gorm.DB // GORM实例

	DBList map[string]*gorm.DB // 多数据源实例
)

type gormPlugin struct{}

func CreateGormPlug() *gormPlugin {
	return &gormPlugin{}
}

func (*gormPlugin) Register(group *gin.RouterGroup) {
	initialize.GormInit()
}

func (*gormPlugin) RouterPath() string {
	return "gorm"
}
