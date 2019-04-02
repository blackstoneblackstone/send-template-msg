package common

import (
	"io/ioutil"
	"log"
	"model"
	"net/http"
)

func HttpGet(url string, jsonModel model.JsonModel) (model.JsonModel, error) {
	//发送请求
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("错误:发送请求", err)
		return nil, err
	}

	//接收到返回数据
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		log.Println("错误:读取body", err)
		return nil, err
	}
	e := jsonModel.JsonToModel(body)
	return jsonModel, e
}
