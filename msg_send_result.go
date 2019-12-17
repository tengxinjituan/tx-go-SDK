package models

// MsgSendResult 短信发送返回结果
type MsgSendResult struct {
	SendResult
}

// EmptyMsgSendResult 返回默认MsgSendResult
func EmptyMsgSendResult() *MsgSendResult {
	return &MsgSendResult{}
}
