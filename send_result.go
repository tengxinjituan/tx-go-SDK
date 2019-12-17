package models

// SendResult 提交返回结果
type SendResult struct {
	MsgId    string `json:"msgId"`    //消息id
	Time     string `json:"time"`     //响应时间
	Code     string `json:"code"`     //状态码
	ErrorMsg string `json:"errorMsg"` //状态码说明（成功返回空）
}
