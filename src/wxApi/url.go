package wxApi

import "fmt"

/*
get access token url
https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s
*/
func AccessTokenUrl(appId string, appSec string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSec)
}
func SimaAccessTokenUrl(appId string) string {
	return fmt.Sprintf("http://data.simamedia.cn/index.php?g=Restful&m=wx&a=access_token&appid=%s", appId);
}
func SendTplMsgUrl(accessToken string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken);
}

func GetOpenidsUrl(accessToken string, nextOpenidString string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s", accessToken, nextOpenidString);
}
