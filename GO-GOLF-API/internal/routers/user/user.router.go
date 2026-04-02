package user

import (
	"GO-GOLF-API/internal/middlewares"
	"GO-GOLF-API/internal/wire"
	"GO-GOLF-API/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandler()
	userService, _ := wire.InitUserServiceHandler()

	// public router
	userRouterPublic := Router.Group("/app/public/user")
	{
		userRouterPublic.GET("/register")
		userRouterPublic.POST("/login", response.Wrap(userController.Login))
		userRouterPublic.POST("/refresh", response.Wrap(userController.Refresh))
		userRouterPublic.GET("/otp")
	}
	// private router
	userRouterPrivate := Router.Group("/app/user")
	// userRouterPrivate.Use(Limiter())
	userRouterPrivate.Use(middlewares.AuthenMiddleware(userService))
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/me", response.Wrap(userController.GetUserById))
		userRouterPrivate.GET("/logout", response.Wrap(userController.LogOut))
	}
}
