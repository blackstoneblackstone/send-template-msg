package common

import (
	"fmt"
	"io"
	"io/ioutil"
	"model"
	"os"
)

type AccessTokenUtil struct {
	AppId, AppSec string
}

func (accessToken *AccessTokenUtil) GetAccessToken() (string, bool) {
	var f *os.File
	var err1 error
	fileName := "/Users/lihongbin/Desktop/life/access_token_" + accessToken.AppId + ".json"
	fileUtil := FileUtil{}
	at := model.AccessToken{}
	if fileUtil.CheckFileIsExist(fileName) { //如果文件存在
		fmt.Println("access token file have exist!")
		read, _ := ioutil.ReadFile(fileName)
		at.JsonToModel(read)
	} else {
		f, err1 = os.Create(fileName) //创建文件
		fmt.Println("access token file does not exist!")
		//写入文件(字符串)
		fileUtil.Check(err1)
		httpUtil := HttpUtil{}
		httpUtil.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx4069e1635ae1be38&secret=4578c042ea9361b6e16626f1aa3d7e52", &at)
		if at.Errcode != 0 {
			return at.Errmsg, false
		} else {
			io.WriteString(f, at.ModelToJson())
		}
	}
	return at.Access_token, true
}
