package wxApi

import (
	"common"
	"encoding/json"
	"log"
)

type Fans struct {
	Errcode      int
	Errmsg       string
	Total, Count int
	Data         Data
	Next_openid  string
}

func (fans *Fans) Refresh(appId string, appSec string, nextOpenid string) {
	fans.getFans(appId, appSec, nextOpenid)
}

func (fans *Fans) getFans(appId string, appSec string, nextOpenid string) {
	accessToken := GetAccessToken(appId, appSec, false)
	log.Println("accessToken->" + accessToken)
	common.HttpGet(GetOpenidsUrl(accessToken, nextOpenid), fans)
	//token expired
	if fans.Errcode == 42001 {
		GetAccessToken(appId, appSec, true)
		fans.getFans(appId, appSec, nextOpenid)
	}
}

type Data struct {
	Openid []string
}

func (fans *Fans) JsonToModel(body []byte) error {
	err := json.Unmarshal(body, &fans)
	return err
}

func (fans *Fans) ModelToJson() string {
	access, _ := json.Marshal(&fans)
	return string(access)
}
