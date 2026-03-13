package initialize

import (
	"GO-ECOMMERCE-BACKEND-API/global"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	var m = global.Config.MySql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.UserName, m.PassWord, m.Host, m.Port, m.DbName)
	global.Logger.Info("InitMysql", zap.String("dsn", s))
	db, err := sqlx.Connect("mysql", s)
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing Mysql Successfully")
	global.Mdb = db
}

func SetPool() {
	var m = global.Config.MySql
	sqlDb := global.Mdb.DB
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}
