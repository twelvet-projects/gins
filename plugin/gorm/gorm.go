package gorm

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvet-s/gins/plugin/gorm/config"
	globalPugin "github.com/twelvet-s/gins/plugin/gorm/global"
	"github.com/twelvet-s/gins/plugin/gorm/initialize"
	"gorm.io/gorm"
)

var (
	INSTANCE_DB *gorm.DB // GORM实例

	INSTANCE_DBList map[string]*gorm.DB // 多数据源实例
)

type gormPlugin struct{}

func CreateGormPlug(config *config.Config) *gormPlugin {
	globalPugin.CONFIG = config
	return &gormPlugin{}
}

func (*gormPlugin) Register(group *gin.RouterGroup) {
	initialize.GormInit()
}

func (*gormPlugin) RouterPath() string {
	return "gorm"
}
