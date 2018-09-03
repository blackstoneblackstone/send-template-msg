package wxApi

import (
	"fmt"
	"io"
	"io/ioutil"
	"model"
	"os"
	"common"
)

func GetAccessToken(AppId, AppSec string) (string) {
	config := common.GetConfig()
	var f *os.File
	var err1 error
	fileName := config.Wechat.AccessTokenPath + "/access_token_" + AppId + ".json"
	fileUtil := common.FileUtil{}
	at := model.AccessToken{}
	if fileUtil.CheckFileIsExist(fileName) { //如果文件存在
		fmt.Println("access token file have exist!")
		read, _ := ioutil.ReadFile(fileName)
		at.JsonToModel(read)
		if (at.Access_token != "") {
			return at.Access_token
		}
	}
	f, err1 = os.Create(fileName) //创建文件
	fmt.Println("access token file does not exist!")
	//写入文件(字符串)
	fileUtil.Check(err1)
	url := SimaAccessTokenUrl(AppId)
	fmt.Println("获取accessToken url->" + url)
	common.HttpGet(url, &at)
	fmt.Println("获取accessToken ->" + at.Access_token)

	if at.Errcode != 0 {
		panic(at.Errmsg)
	} else {
		io.WriteString(f, at.ModelToJson())
	}
	return at.Access_token
}
