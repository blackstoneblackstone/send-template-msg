package main

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"common"
	"dbServer"
)

func main() {

	//配置文件初始化
	config := common.GetConfig()
	//redisApi := dbServer.RedisApi{Config: config}
	//redisApi.Connect()
	//s, _ := redisApi.Get("cool")
	//fmt.Println(s)

	mysqlApi := dbServer.CreateMysqlApi(config)
	appSec, _ := mysqlApi.GetWxApp("wx4069e1635ae1be38")
	fmt.Print(appSec)
}
