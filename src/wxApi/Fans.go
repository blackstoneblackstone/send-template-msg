package wxApi

import (
	"common"
	"encoding/json"
	"fmt"
)

type Fans struct {
	Total, Count int
	Data         Data
	Next_openid string
}

func (fans *Fans) Refresh(appId string, appSec string, nextOpenid string) {
	accessToken := GetAccessToken(appId, appSec)
	fmt.Println("accessToken->"+accessToken)
	common.HttpGet(GetOpenidsUrl(accessToken, nextOpenid), fans)
}

type Data struct {
	Openid      []string
}

func (fans *Fans) JsonToModel(body []byte) error {
	err := json.Unmarshal(body, &fans)
	return err
}

func (fans *Fans) ModelToJson() string {
	access, _ := json.Marshal(&fans)
	return string(access)
}
