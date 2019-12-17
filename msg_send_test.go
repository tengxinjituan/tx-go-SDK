package apis

import (
	"fmt"
	"testing"
)

// Test_MsgSend 短信发送测试
func Test_MsgSend(t *testing.T) {
	batchSend := NewMsgSend("10.10.50.3:8082", "lijie01", "host:62115", "")
	// {
	// 	"account":"520520",
	// 	"password":"520520_520520",
	// 	"msg":"验证码：mtzg",
	// 	"phone":"18732642929",
	// 	"sendtime":"",
	// 	"report":"true",
	// 	"extend":"20",
	// 	"uid":"1000099",
	// 	"format":"",
	// 	"useragent":""
	// }
	amodel := batchSend.NewModel("朱书彦00000979", "13633861512", "", "", "1000102")
	aresult, aerr := batchSend.Send(amodel)
	t.Log(aresult, aerr)
	fmt.Println(aresult, aerr)
}

// Test_MsgSendWithRsa rsa加密方式短信发送测试
func Test_MsgSendWithRsa(t *testing.T) {
	batchSend := NewMsgSend("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)
	// {
	// 	"account":"lijie01",
	// 	"password":"host:62115",
	// 	"msg":"您的验证码：909090【世界和平】",
	// 	"phone":"18703807708",
	// 	"sendtime":"201903271550",
	// 	"report":"true",
	// 	"extend":"555",
	// 	"uid":"1000102",
	// 	"format":"json",
	// 	"useragent":"http"
	// }
	amodel := batchSend.NewModel("朱书彦00000979", "13633861512", "", "", "1000102")
	aresult, aerr := batchSend.SendWithRsa(amodel)
	t.Log(aresult, aerr)
	fmt.Println(aresult, aerr)
}
