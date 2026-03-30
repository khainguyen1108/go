package initialize

import (
	"GO-ECOMMERCE-BACKEND-API/global"
	"GO-ECOMMERCE-BACKEND-API/internal/routers"
	"GO-ECOMMERCE-BACKEND-API/pkg/constant"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	switch global.Config.Server.Mode {
	case constant.DEV:
		{
			gin.SetMode(gin.DebugMode)
			gin.ForceConsoleColor()
			r = gin.Default()
		}
	case constant.PROD:
		{
			gin.SetMode(gin.ReleaseMode)
			r = gin.New()
		}
	default:
		{
			gin.SetMode(gin.ReleaseMode)
			r = gin.New()
		}
	}

	// r.Use() // logging
	// r.Use() // cross
	// r.Use() //limiter
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus") // tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
	}

	return r
}
