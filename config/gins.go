package config

import (
	gormConfig "github.com/twelvet-s/gins/plugin/gorm/config"
	redisConfig "github.com/twelvet-s/gins/plugin/redis/config"
)

type Gins struct {
	Gorm gormConfig.Config // Gorm配置

	Redis redisConfig.Redis // redis配置
	Zap   Zap               `mapstructure:"zap" json:"zap" yaml:"zap"` // 日志配置
}
