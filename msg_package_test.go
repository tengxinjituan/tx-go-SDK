package apis

import (
	"fmt"
	"testing"
)

// Benchmark_MsgPackage 短信包测试
func Benchmark_MsgPackage(b *testing.B) {
	packageSend := NewMsgPackage("10.10.50.3:8082", "lijie01", "host:62115", "")
	//account=N6000001&password=123456&msg=132769909863|测试短信1&msg=13717536594|测试短信2&msg=15913912569|测试短信3&sendtime=201704101400&report=true&extend=555&uid=321abc&format=json&useragent=http
	cmodel := packageSend.NewModel("msg=132769909863|测试短信1&msg=13717536594|测试短信2&msg=15913912569|测试短信3", "", "", "1000099")
	cresult, cerr := packageSend.Send(cmodel)
	b.Log(cresult, cerr)
	fmt.Println(cresult, cerr)
}

// Test_MsgPackageRSA 短信包加密方式测试
func Test_MsgPackageRSA(t *testing.T) {
	packageSend := NewMsgPackage("10.10.50.3:8082", "lijie01", "host:62115", TestPublicKey)
	//account=N6000001&password=123456&msg=132769909863|测试短信1&msg=13717536594|测试短信2&msg=15913912569|测试短信3&sendtime=201704101400&report=true&extend=555&uid=321abc&format=json&useragent=http
	cmodel := packageSend.NewModel("msg=132769909863|测试短信1&msg=13717536594|测试短信2&msg=15913912569|测试短信3", "", "", "1000099")
	cresult, cerr := packageSend.SendRSA(cmodel)
	t.Log(cresult, cerr)
	fmt.Println(cresult, cerr)
}
