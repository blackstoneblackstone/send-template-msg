package wxApi

import "fmt"

/*
get access token url
*/
func AccessTokenUrl(appId string, appSec string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSec)
}

func SendTplMsgUrl(accessToken string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken);
}
