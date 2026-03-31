package initialize

import (
	"GO-GOLF-API/global"
	"GO-GOLF-API/internal/middlewares"
	"GO-GOLF-API/internal/routers"
	"GO-GOLF-API/pkg/constant"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	switch global.Config.Server.Mode {
	case constant.DEV:
		{
			gin.SetMode(gin.DebugMode)
			gin.ForceConsoleColor()
			r = gin.New()
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
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CustomRecoveryMiddleware())
	r.Use(middlewares.ErrorHandlingMiddleware())
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
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
