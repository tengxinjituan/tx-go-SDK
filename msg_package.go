package models

import "fmt"

// MsgPackage 包短信
type MsgPackage struct {
	Send
}

// NewMsgPackage 初始化包发送实体
func NewMsgPackage(account, password, msg, sendtime, extend, uid string) *MsgPackage {
	result := &MsgPackage{}

	result.Account = account
	result.Password = password
	result.Msg = msg
	result.Report = true
	result.Extend = extend
	result.UId = uid
	result.Format = "json"
	result.UserAgent = "http"

	return result
}

// ToURLEncode 转换为key=Value&key=Value格式
func (s *MsgPackage) ToURLEncode() string {
	// params := url.Values{}
	// params.Add("account", s.Account)
	// params.Add("password", s.Password)
	// params.Add("extend", s.Extend)
	// params.Add("uid", s.UId)
	// params.Add("useragent", s.UserAgent)
	// result := params.Encode()
	// result += fmt.Sprintf("&report=%t", s.Report)
	// result += "&" + s.Msg
	// return result
	return fmt.Sprintf("account=%s&password=%s&extend=%s&format=%s&%s&sendtime=%s&uid=%s&useragent=%s&report=%t",
		s.Account, s.Password, s.Extend, s.Format, s.Msg, s.SendTime, s.UId, s.UserAgent, s.Report)
}
