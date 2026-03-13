package initialize

import (
	"GO-ECOMMERCE-BACKEND-API/global"
	StringUtils "GO-ECOMMERCE-BACKEND-API/pkg/utils"
	"fmt"
)

func Run() {
	//load config enviroment
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.MySql.UserName)
	//init logger
	InitLogger()
	//connect db
	InitMysql()
	//connect redis
	InitRedis()
	r := InitRouter()
	r.Run(StringUtils.GetServerPort(global.Config.Server.Port))
}
