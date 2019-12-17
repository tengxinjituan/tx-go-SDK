package models

// MsgSend 批量提交短信
type MsgSend struct {
	Send
	Phone string `json:"phone"` //手机号码。多个手机号码使用英文逗号分隔
}

// NewMsgSend 返回批量提交实体
func NewMsgSend(account, password, msg, phone, sendtime, extend, uid string) *MsgSend {
	result := &MsgSend{
		Phone: phone,
	}

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

// EmptyMsgSend 初始化默认MsgSend
func EmptyMsgSend() *MsgSend {
	return &MsgSend{}
}
