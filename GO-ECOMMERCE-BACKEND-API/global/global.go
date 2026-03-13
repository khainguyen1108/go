package global

import (
	"GO-ECOMMERCE-BACKEND-API/pkg/logger"
	"GO-ECOMMERCE-BACKEND-API/pkg/setting"

	"github.com/jmoiron/sqlx"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *sqlx.DB
)
