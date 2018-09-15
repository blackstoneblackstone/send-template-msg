package main

import (
	"wxApi"
	"os"
	"log"
	"dbServer"
	"time"
)

func main() {
	//./bin/fans wx3be7b35d2d7a8256 ""
	//配置文件初始化"wx293dbb0f011bcac3"
	//env GOOS=linux GOARCH=amd64 go build src/fans.go
	appId := os.Args[1]
	nextO := os.Args[2]
	liqud := 5000
	mysqlApi := dbServer.CreateMysqlApi()
	appSec, _ := mysqlApi.GetWxApp(appId)
	fans := wxApi.Fans{}
	fans.Refresh(appId, appSec, nextO)
	// openid total
	total := fans.Total
	if fans.Errcode != 0 {
		log.Println(fans.Errmsg)
	}
	page := total/10000 + 1
	log.Printf("openid length is %d", total)
	openIds := fans.Data.Openid
	count := 0
	for page > 0 {
		for i := 0; i < (len(openIds)/liqud + 1); i++ {
			t := liqud * (i + 1)
			if liqud*(i+1) > len(openIds) {
				t = len(openIds)
				count = count + len(openIds)%liqud
			} else {
				count = count + liqud
			}
			log.Printf("count is -> %d", count)
			o := openIds[liqud*i : t]
			mysqlApi.SaveOpenIds(appId, o)
		}
		page--
		openIds = openIdEx(&fans, appId, appSec, fans.Next_openid)
	}
	log.Println("program have finish .....")
	i := 5
	for i < 6 && i > 0 {
		i--
		log.Printf("%d .....", i)
		time.Sleep(time.Second)
	}
}

func openIdEx(fans *wxApi.Fans, appId string, appSec string, nextOpenId string) []string {
	fans.Refresh(appId, appSec, nextOpenId)
	return fans.Data.Openid
}
