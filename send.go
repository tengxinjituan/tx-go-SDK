package models

// Send 短信提交发送
type Send struct {
	Account   string `json:"account"`   //用户账号
	Password  string `json:"password"`  //用户密码
	Msg       string `json:"msg"`       //短信内容。长度不能超过536个字符
	SendTime  string `json:"sendtime"`  //定时发送短信时间。格式为yyyyMMddHHmm，值小于或等于当前时间则立即发送，默认立即发送，选填
	Report    bool   `json:"report"`    //是否需要状态报告（默认false），选填
	Extend    string `json:"extend"`    //下发短信号码扩展码，纯数字，建议1-3位，选填
	UId       string `json:"uid"`       //该条短信在您业务系统内的ID，如产品号或者短信发送记录流水号，选填
	Format    string `json:"format"`    //请求相应格式json或者xml或者txt，选填
	UserAgent string `json:"useragent"` //来源(cmpp,web,http)，默认http，选填
}
