package models

import (
	"net/url"
)

// RsaSend rsa加密发送包体
type RsaSend struct {
	Key     string `json:"key"`     //AES-key(进行RSA加密)
	Data    string `json:"data"`    //密文(对数据原文使用AES-KEY进行AES加密)
	Account string `json:"account"` //账号
}

// ToURLEncode 格式为key=value&key=value字符串格式
func (s *RsaSend) ToURLEncode() string {
	params := url.Values{}
	params.Add("key", s.Key)
	params.Add("data", s.Data)
	params.Add("account", s.Account)
	return params.Encode()
}
