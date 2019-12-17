package models

// MsgVariable 变量短信发送
type MsgVariable struct {
	Send
	Params string `json:"params"` //手机号码和变量参数，多组参数使用英文分号;区分，必填；字段最多不能超过1000个参数组
}

// NewMsgVariable 生成变量短信提交包
func NewMsgVariable(account, password, msg, params, sendtime, extend, uid string) *MsgVariable {
	result := &MsgVariable{
		Params: params,
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
