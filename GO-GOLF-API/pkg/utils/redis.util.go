package utils

import (
	"GO-GOLF-API/global"
	"context"
	"time"
)

func Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return global.Rdb.Set(ctx, key, value, expiration).Err()
}

func Get(ctx context.Context, key string) (interface{}, error) {
	return global.Rdb.Get(ctx, key).Result()
}

func Delete(ctx context.Context, keys ...string) error {
	return global.Rdb.Del(ctx, keys...).Err()
}
