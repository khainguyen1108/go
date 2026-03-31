package user

import (
	"GO-GOLF-API/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandler()

	// public router
	userRouterPublic := Router.Group("/app/public/user")
	{
		userRouterPublic.GET("/register", userController.GetUserById)
		userRouterPublic.POST("/login", userController.Login)
		userRouterPublic.GET("/otp")
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
