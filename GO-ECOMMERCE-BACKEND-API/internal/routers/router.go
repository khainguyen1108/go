package routers

import (
	"GO-ECOMMERCE-BACKEND-API/internal/container"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", container.TestControllerInstance.Pong)
		v1.GET("/user/1", container.UserControllerInstance.GetUserById)
	}

	return r
}
