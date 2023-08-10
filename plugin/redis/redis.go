package redis

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/twelvet-s/gins/framework/global"
	"github.com/twelvet-s/gins/plugin/redis/config"
	globalPlugin "github.com/twelvet-s/gins/plugin/redis/global"
	"go.uber.org/zap"
)

var (
	INSTANCE *redis.Client // redis实例
)

type PluginRedis struct{}

func CreateRedisPlug(redisConfig *config.Redis) *PluginRedis {
	globalPlugin.CONFIG = redisConfig
	return &PluginRedis{}
}

func (*PluginRedis) Register(group *gin.RouterGroup) {
	client := redis.NewClient(&redis.Options{
		Addr:     globalPlugin.CONFIG.Addr,
		Password: globalPlugin.CONFIG.Password, // no password set
		DB:       globalPlugin.CONFIG.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		INSTANCE = client
	}
}

func (*PluginRedis) RouterPath() string {
	return "redis"
}
