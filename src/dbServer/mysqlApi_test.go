package dbServer

import (
	"testing"
	"common"
)

func TestMysqlApi_GetWxApp(t *testing.T) {
	tests := []struct{ appId, appSec string }{
		{"wx3be7b35d2d7a8256", "2"},
		{"wx293dbb0f011bcac3", "3"},
	}
	config := common.GetConfig()
	mysqlApi := CreateMysqlApi(config)
	for _, tt := range tests {
		if appSec, _ := mysqlApi.GetWxApp(tt.appId); appSec != tt.appSec {
			t.Errorf("select mysql by appid get app secret is error")
		}
	}
}
