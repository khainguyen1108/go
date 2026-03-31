package global

import (
	"GO-GOLF-API/pkg/logger"
	"GO-GOLF-API/pkg/setting"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *sqlx.DB
	Rdb    *redis.Client
)
