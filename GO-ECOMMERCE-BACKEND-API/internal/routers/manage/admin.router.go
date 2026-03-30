package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("/admin")
	{
		userRouterPublic.POST("/login")
	}
	// private router
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/active_user")
	}
}
