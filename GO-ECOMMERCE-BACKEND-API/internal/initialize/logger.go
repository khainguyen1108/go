package initialize

import (
	"GO-ECOMMERCE-BACKEND-API/global"
	"GO-ECOMMERCE-BACKEND-API/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Log)
}
