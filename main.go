package main

import (
	"fmt"

	"github.com/zsy619/sms870/apis"
)

func main() {
	sms := apis.NewSMS("10.10.50.3:8082", "lijie01", "host:62115", apis.TestPublicKey)

	amodel := sms.NewMsgSendModel("朱书彦00000979", "13633861512", "", "", "1000102")
	aresult, aerr := sms.MsgSend(amodel)
	fmt.Println(aresult, aerr)

	fmt.Println(sms)
}
