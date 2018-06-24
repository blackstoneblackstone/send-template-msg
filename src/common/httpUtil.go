package common

import (
	"fmt"
	"io/ioutil"
	"model"
	"net/http"
)

type HttpUtil struct {
}

func (httpUtil HttpUtil) Get(url string, jsonModel model.JsonModel) (model.JsonModel, error) {
	//发送请求
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("错误:发送请求", err)
		return nil, err
	}

	//接收到返回数据
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		fmt.Println("错误:读取body", err)
		return nil, err
	}
	e := jsonModel.JsonToModel(body)
	return jsonModel, e
}
