package middlewares

import (
	"GO-ECOMMERCE-BACKEND-API/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrCodeParamInvalid)
			c.Abort()
			return
		}

		c.Next()
	}
}
