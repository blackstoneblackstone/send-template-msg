package main

import (
	"wxApi"
	"os"
	"log"
	"dbServer"
	"strconv"
)

func main() {

	//配置文件初始化"wx293dbb0f011bcac3"
	//nv GOOS=linux GOARCH=amd64 go build src/main.go
	appId := os.Args[1]
	start, _ := strconv.Atoi(os.Args[2])
	mysqlApi := dbServer.CreateMysqlApi()
	appSec, _ := mysqlApi.GetWxApp(appId)
	//fmt.Print(appSec)
	//y := "afsaafafasfafa"
	//x := "sfa"
	//r := Count(y, x)
	//fmt.Print(r)
	fans := wxApi.Fans{}
	fans.Refresh(appId, appSec, "")
	// openid total
	total := fans.Total
	page := total / 10000
	c := make(chan int, 1)
	log.Printf("openid length is %d", total)
	openIds := fans.Data.Openid
	count := 0
	for page > 0 {
		for i := start; i < len(openIds); i++ {
			count++
			c <- count
			go mysqlApi.SaveOpenIds(appId, openIds[i], c)
		}
		openIds = openIdEx(&fans, appId, appSec, fans.Next_openid)
	}
}

func openIdEx(fans *wxApi.Fans, appId string, appSec string, nextOpenId string) []string {
	fans.Refresh(appId, appSec, nextOpenId)
	return fans.Data.Openid
}
