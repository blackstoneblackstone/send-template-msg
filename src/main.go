package src

import (
	"fmt"
	"common"
	"dbServer"
)

func main() {

	//配置文件初始化

	mysqlApi := dbServer.CreateMysqlApi()
	appSec, _ := mysqlApi.GetWxApp("wx4069e1635ae1be38")
	fmt.Print(appSec)
}
