package wxApi

import (
	"common"
	"model"
	"fmt"
)

type TemplateMsgApi struct {
}

func (templateMsgApi *TemplateMsgApi) sendMsg(msg model.SendMsg) {
	accessToken := GetAccessToken(msg.AppId, msg.AppSec, false)
	tplMsg, _ := common.HttpGet(SendTplMsgUrl(accessToken), &msg)
	fmt.Print(tplMsg)
}
