package wxApi

import (
	"bufio"
	"common"
	"dbServer"
	"fmt"
	"io"
	"log"
	"os"
)

func GetOpenIdFromText(appId string, txtPath string) {
	mySqlApi := dbServer.CreateMysqlApi()
	mySqlApi.DeleteAppIdBySingle(appId)
	fileName := txtPath
	fileUtil := common.FileUtil{}
	if fileUtil.CheckFileIsExist(fileName) {
		fi, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer fi.Close()
		br := bufio.NewReader(fi)
		var openIds []string
		count := 0
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				//最后save下
				mySqlApi.SaveOpenIdsBySingle(appId, openIds)
				log.Printf("insert into count -> %d", count)
				break
			}
			count++
			openIds = append(openIds, string(a))
			if count%100 == 0 {
				mySqlApi.SaveOpenIdsBySingle(appId, openIds)
				openIds = nil
				log.Printf("insert into count -> %d", count)
			}
		}
	}
}
