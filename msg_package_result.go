package models

// MsgPackageResult 包短信发送返回结果
type MsgPackageResult struct {
	SendResult
	FailNum int `json:"failNum"` //失败号码数
	SuccNum int `json:"succNum"` //成功号码数
}

// EmptyMsgPackageResult 默认包短信返回结果
func EmptyMsgPackageResult() *MsgPackageResult {
	return &MsgPackageResult{}
}
