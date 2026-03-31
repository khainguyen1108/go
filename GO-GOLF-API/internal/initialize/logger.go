package initialize

import (
	"GO-GOLF-API/global"
	"GO-GOLF-API/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Log)
}
