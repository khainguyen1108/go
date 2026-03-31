package middlewares

import (
	"time"

	"GO-GOLF-API/global" // Import biến global chứa Zap Logger

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("STATUS:", statusCode),
			zap.String("METHOD:", c.Request.Method),
			zap.String("PATH:", c.Request.RequestURI),
			zap.Time("START_TIME:", startTime),
			zap.Duration("LATENCY:", latency),
			zap.String("CLIENT_IP:", c.ClientIP()),
			zap.String("USER_AGENT:", c.Request.UserAgent()),
		}

		global.Logger.Info("[HTTP Access]", logFields...)
	}
}
