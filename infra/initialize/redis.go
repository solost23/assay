package initialize

import (
	"assay/infra/global"
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const defaultDB = 0

func initRedis() {
	redisConfig := global.ServerConfig.Redis
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       defaultDB,
	})

	if _, err := global.RDB.Ping(context.TODO()).Result(); err != nil {
		zap.S().Panic("redis connect failed: ", err)
	}
}
