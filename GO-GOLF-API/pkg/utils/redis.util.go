package utils

import (
	"GO-GOLF-API/global"
	"context"
	"time"
)

func Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return global.Rdb.Set(ctx, key, value, expiration).Err()
}

func GetInt64(ctx context.Context, key string) (int64, error) {
	return global.Rdb.Get(ctx, key).Int64()
}

func Delete(ctx context.Context, keys ...string) error {
	return global.Rdb.Del(ctx, keys...).Err()
}
