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
