package wxApi

import (
	"common"
	"fmt"
	"model"
)

type TemplateMsgApi struct {
}

func (templateMsgApi *TemplateMsgApi) sendMsg(msg model.SendMsg) {
	accessToken := GetAccessToken(msg.AppId, msg.AppSec, false)
	tplMsg, _ := common.HttpGet(SendTplMsgUrl(accessToken), &msg)
	fmt.Print(tplMsg)
}
