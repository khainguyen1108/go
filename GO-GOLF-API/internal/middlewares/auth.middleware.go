package middlewares

import (
	"GO-GOLF-API/internal/service"
	"GO-GOLF-API/pkg/response"
	JwtUtil "GO-GOLF-API/pkg/utils"
	RedisUtils "GO-GOLF-API/pkg/utils"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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
		userId := int(claims["user_id"].(float64))
		iat := int64(claims["iat"].(float64))
		timeLimit, errRedis := RedisUtils.Get(context.Background(), strconv.Itoa(userId))
		if errRedis != nil && errRedis != redis.Nil {
			response.HandleGlobalError(c, errRedis)
			c.Abort()
			return
		}

		if errRedis != redis.Nil {
			timeStr, ok := timeLimit.(string)
			if !ok {
				response.HandleGlobalError(c, response.NewAppError(http.StatusInternalServerError, response.ErrInternalError, gin.Error{}))
			}
			ts, errParse := strconv.ParseInt(timeStr, 10, 64)
			if errParse != nil {
				response.HandleGlobalError(c, response.NewAppError(http.StatusInternalServerError, response.ErrInternalError, errParse))
			}
			if ts > iat {
				response.HandleGlobalError(c, response.NewAppError(http.StatusUnauthorized, response.ErrUserAlreadyBlocked, gin.Error{}))
				c.Abort()
				return
			}
		}

		//get user info
		userInfo, userErr := userService.GetUserInfoById(userId)

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
