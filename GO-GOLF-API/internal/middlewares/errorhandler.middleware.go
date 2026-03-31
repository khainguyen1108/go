package middlewares

import (
	"GO-GOLF-API/global"
	"GO-GOLF-API/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*response.AppError); ok {
				global.Logger.Error("Request Error", zap.String("path", c.Request.URL.Path), zap.Error(err))
				response.ErrorResponse(c, appErr.StatusCode, appErr.AppCode)
			} else {
				global.Logger.Error("System Error", zap.Error(err))
				response.ErrorResponse(c, http.StatusInternalServerError, response.ErrInternalError)
			}
		}
	}
}
