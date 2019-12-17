package apis

import (
	"fmt"
	"testing"
)

// Test_SMSMsgSend 短信发送测试
func Test_SMSMsgSend(t *testing.T) {
	sms := NewSMS("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)

	model := sms.NewMsgSendModel("朱书彦00000979", "13633861512", "", "", "1000102")
	result, err := sms.MsgSend(model)
	fmt.Println(result, err)
	fmt.Println(sms)
}

// Test_SMSMsgSendWithRsa rsa加密方式短信发送测试
func Test_SMSMsgSendWithRsa(t *testing.T) {
	sms := NewSMS("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)

	model := sms.NewMsgSendModel("朱书彦00000979", "13633861512", "", "", "1000102")
	result, err := sms.MsgSendWithRsa(model)
	fmt.Println(result, err)
	fmt.Println(sms)
}

// Test_SMSMsgPackage 短信包测试
func Test_SMSMsgPackage(t *testing.T) {
	sms := NewSMS("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)
	//account=N6000001&password=123456&msg=132769909863|测试短信1&msg=13717536594|测试短信2&msg=15913912569|测试短信3&sendtime=201704101400&report=true&extend=555&uid=321abc&format=json&useragent=http
	cmodel := sms.NewMsgPackageModel("msg=132769909863|测试短信1&msg=13717536594|测试短信2&msg=15913912569|测试短信3", "", "", "1000099")
	cresult, cerr := sms.MsgPackage(cmodel)
	t.Log(cresult, cerr)
	fmt.Println(cresult, cerr)
}

// Test_SMSMsgPackageWithRsa 短信包加密方式测试
func Test_SMSMsgPackageWithRsa(t *testing.T) {
	sms := NewSMS("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)
	cmodel := sms.NewMsgPackageModel("msg=132769909863|测试短信1&msg=13717536594|测试短信2&msg=15913912569|测试短信3", "", "", "1000099")
	cresult, cerr := sms.MsgPackageWithRsa(cmodel)
	t.Log(cresult, cerr)
	fmt.Println(cresult, cerr)
}

// Test_SMSMsgVariable 变量短信测试
func Test_SMSMsgVariable(t *testing.T) {
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
	sms := NewSMS("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)
	bmodel := sms.NewMsgVariableModel("验证码短信{$var}你好，{$var}请你于{$var}日参加考试", "18697599572,张三1,男,19;13155370661,张三2,女,20", "", "", "1000102")
	bresult, berr := sms.MsgVariable(bmodel)
	t.Log(bresult, berr)
	fmt.Println(bresult, berr)
}

// Test_SMSMsgVariableWithRsa 变量短信测试
func Test_SMSMsgVariableWithRsa(t *testing.T) {
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
	sms := NewSMS("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)
	bmodel := sms.NewMsgVariableModel("验证码短信{$var}你好，{$var}请你于{$var}日参加考试", "18697599572,张三1,男,19;13155370661,张三2,女,20", "", "", "1000102")
	bresult, berr := sms.MsgVariableWithRsa(bmodel)
	t.Log(bresult, berr)
	fmt.Println(bresult, berr)
}
