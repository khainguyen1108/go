package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func ValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validate := validator.New()

		c.Set("validation", validate)

		c.Next()
	}
}
