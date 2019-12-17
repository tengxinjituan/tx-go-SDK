package apis

import (
	"fmt"
	"testing"
)

// Test_VariantSend 变量短信测试
func Test_VariantSend(t *testing.T) {
	variableSend := NewMsgVariable("10.10.50.3:8082", "lijie01", "host:62115", "")
	// {
	//     "account":"admin_lh",
	//     "password":"123asd",
	//     "msg":"【腾信集团1】验证码短信{$var}你好，{$var}请你于{$var}日参加考试",
	//     "params":"18697599572,张三1,男,19;13155370661,张三2,女,20",
	//     "sendtime":"",
	//     "report":"true",
	//     "extend":"111",
	//     "uid":"1000000",
	//     "format":"",
	//     "useragent": "variable"
	// }
	bmodel := variableSend.NewModel("验证码短信{$var}你好，{$var}请你于{$var}日参加考试", "18697599572,张三1,男,19;13155370661,张三2,女,20", "", "", "1000102")
	bresult, berr := variableSend.Send(bmodel)
	t.Log(bresult, berr)
	fmt.Println(bresult, berr)
}

// Test_MsgVariantWithRsa 变量短信测试
func Test_MsgVariantWithRsa(t *testing.T) {
	variableSend := NewMsgVariable("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)
	// {
	//     "account":"admin_lh",
	//     "password":"123asd",
	//     "msg":"【腾信集团1】验证码短信{$var}你好，{$var}请你于{$var}日参加考试",
	//     "params":"18697599572,张三1,男,19;13155370661,张三2,女,20",
	//     "sendtime":"",
	//     "report":"true",
	//     "extend":"111",
	//     "uid":"1000000",
	//     "format":"",
	//     "useragent": "variable"
	// }
	bmodel := variableSend.NewModel("验证码短信{$var}你好，{$var}请你于{$var}日参加考试", "18697599572,张三1,男,19;13155370661,张三2,女,20", "", "", "1000102")
	bresult, berr := variableSend.SendWithRsa(bmodel)
	t.Log(bresult, berr)
	fmt.Println(bresult, berr)
}
