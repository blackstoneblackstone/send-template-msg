package model

import (
	"encoding/json"
)

type JsonModel interface {
	JsonToModel(body []byte) error
	ModelToJson() string
}

//{"access_token":"ACCESS_TOKEN","expires_in":7200}
type AccessToken struct {
	Access_token string
	Expires_in   int
	Errcode      int
	Errmsg       string
}

func (accessToken *AccessToken) JsonToModel(body []byte) error {
	err := json.Unmarshal(body, &accessToken)
	return err
}

func (accessToken *AccessToken) ModelToJson() string {
	access, _ := json.Marshal(&accessToken)
	return string(access)
}

/*
  "touser": openId,
  "template_id": templateId,
  "url": tplUrl,
  "data": data
*/

type SendMsg struct {
	ToUser, TemplateId, Url, Data, AppId, AppSec string
}

func (sendMsg *SendMsg) ModelToJson() string {
	access, _ := json.Marshal(&sendMsg)
	return string(access)
}

func (sendMsg *SendMsg) JsonToModel(body []byte) error {
	err := json.Unmarshal(body, &sendMsg)
	return err
}
