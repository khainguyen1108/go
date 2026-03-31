package initialize

import (
	"GO-GOLF-API/global"
	StringUtils "GO-GOLF-API/pkg/utils"
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
