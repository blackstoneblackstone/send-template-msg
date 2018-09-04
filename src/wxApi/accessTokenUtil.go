package wxApi

import (
	"fmt"
	"io"
	"io/ioutil"
	"model"
	"os"
	"common"
	"log"
)

//isFresh刷新token
func GetAccessToken(AppId, AppSec string, isRefresh bool) (string) {
	config := common.GetConfig()
	var f *os.File
	var err1 error
	fileName := config.Wechat.AccessTokenPath + "/access_token_" + AppId + ".json"
	fileUtil := common.FileUtil{}
	at := model.AccessToken{}
	if fileUtil.CheckFileIsExist(fileName) && !isRefresh { //如果文件存在
		fmt.Println("access token file have exist!")
		read, _ := ioutil.ReadFile(fileName)
		at.JsonToModel(read)
		if (at.Access_token != "") {
			log.Println("FILE获取accessToken ->" + at.Access_token)
			return at.Access_token
		}
	}
	os.Remove(fileName)
	f, err1 = os.Create(fileName) //创建文件
	//写入文件(字符串)
	fileUtil.Check(err1)
	url := SimaAccessTokenUrl(AppId)
	common.HttpGet(url, &at)
	log.Println("URL获取accessToken ->" + at.Access_token)
	if at.Errcode != 0 {
		panic(at.Errmsg)
	} else {
		io.WriteString(f, at.ModelToJson())
	}
	return at.Access_token
}
