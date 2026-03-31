package initialize

import (
	"GO-GOLF-API/global"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Redis initialization Error: %v", zap.Error(err))
	}

	global.Logger.Info("InitRedis is running")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	errorRedisException(err)
	value, err := global.Rdb.Get(ctx, "score").Result()
	errorRedisException(err)

	global.Logger.Info("value score is::", zap.String("score", value))
}

func errorRedisException(err error) {
	if err != nil {
		global.Logger.Error("Error redis setting", zap.Error(err))
		return
	}
}
