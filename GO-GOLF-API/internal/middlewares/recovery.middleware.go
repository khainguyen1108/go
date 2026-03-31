package middlewares

import (
	"GO-GOLF-API/global"
	"GO-GOLF-API/pkg/response"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CustomRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer handlePanic(c)
		c.Next()
	}
}

func handlePanic(c *gin.Context) {
	if err := recover(); err != nil {
		global.Logger.Error("[Exception]",
			zap.Any("error", err),
			zap.String("request_path", c.Request.URL.Path),
			zap.String("stack_trace", string(debug.Stack())),
		)
		response.ErrorResponse(c, http.StatusInternalServerError, response.ErrInternalError)
		c.Abort()
	}
}
