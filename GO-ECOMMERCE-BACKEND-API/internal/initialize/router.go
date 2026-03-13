package initialize

import (
	"GO-ECOMMERCE-BACKEND-API/internal/container"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", container.TestControllerInstance.Pong)
		v1.GET("/ping2", container.TestControllerInstance.PongOpenApi)
		v1.GET("/user/:id", container.UserControllerInstance.GetUserById)
	}

	return r
}
