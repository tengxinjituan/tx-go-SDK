package models

// MsgVariableResult 变量短信发送结果
type MsgVariableResult struct {
	SendResult
	FailNum int `json:"failNum"`
	SuccNum int `json:"succNum"`
}

// EmptyMsgVariableResult 默认变量短信返回结果
func EmptyMsgVariableResult() *MsgVariableResult {
	return &MsgVariableResult{}
}
