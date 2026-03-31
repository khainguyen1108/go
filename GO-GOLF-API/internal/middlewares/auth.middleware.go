package middlewares

import (
	"GO-GOLF-API/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, http.StatusUnauthorized, response.ErrCodeParamInvalid)
			c.Abort()
			return
		}

		c.Next()
	}
}
