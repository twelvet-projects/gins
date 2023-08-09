package redis

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/twelvet-s/gins/framework/global"
	globalPlugin "github.com/twelvet-s/gins/plugin/redis/global"
	"github.com/twelvet-s/gins/plugin/redis/utils"
	"go.uber.org/zap"
)

type redisPlugin struct{}

func CreateRedisPlug(addr, password string, db int) *redisPlugin {
	globalPlugin.CONFIG.Addr = addr
	globalPlugin.CONFIG.DB = db
	globalPlugin.CONFIG.Password = password
	return &redisPlugin{}
}

func (*redisPlugin) Register(group *gin.RouterGroup) {
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
		utils.REDIS = client
	}
}

func (*redisPlugin) RouterPath() string {
	return "redis"
}
