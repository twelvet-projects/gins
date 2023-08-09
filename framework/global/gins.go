package global

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/twelvet-s/gins/config"
	"go.uber.org/zap"
)

var (
	API    *gin.Engine   // 路由配置
	Viper  *viper.Viper  // yml配置
	CONFIG config.Config // 全局应用配置
	LOG    *zap.Logger   // 日志
)
