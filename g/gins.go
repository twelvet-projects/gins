package g

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/twelvet-s/gins/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	API    *gin.Engine   // 路由配置
	DB     *gorm.DB      // 数据库
	REDIS  *redis.Client // reids
	VP     *viper.Viper  // yml配置
	CONFIG config.Config // 全局应用配置
	LOG    *zap.Logger   // 日志
)
