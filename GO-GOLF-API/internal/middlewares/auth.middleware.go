package middlewares

import (
	"GO-GOLF-API/internal/service"
	"GO-GOLF-API/pkg/response"
	JwtUtil "GO-GOLF-API/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware(userService service.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			err := response.NewAppError(http.StatusUnauthorized, response.ErrTokenNotFound, gin.Error{})
			response.HandleGlobalError(c, err)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := JwtUtil.VerifyAndParseJWT(tokenString)

		if err != nil {
			response.HandleGlobalError(c, err)
			c.Abort()
			return
		}
		//get user info
		userInfo, userErr := userService.GetUserInfoById(int(claims["user_id"].(float64)))

		if userErr != nil {
			response.HandleGlobalError(c, userErr)
			c.Abort()
			return
		}

		c.Set("user", userInfo)
		c.Set("section_id", claims["section_id"])

		c.Next()
	}
}
