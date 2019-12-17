package apis

import "strings"

// Sending 短信发送
type Sending struct {
	// URL 提交地址，例如：192.168.0.10:8080，不包括 https:// 与 / 等
	URL        string `json:"url"`      //提交地址
	Account    string `json:"account"`  //账号
	Password   string `json:"password"` //密码
	PublicKey  []byte //RSA公钥
	sendURL    string //发送地址
	sendRSAUrl string //加密发送地址
}

func newSend(url, account, password string) *Sending {
	result := &Sending{}
	result.URL = url
	result.Account = account
	result.Password = password
	// URL地址判定
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		result.URL = "http://" + url
	}
	if !strings.HasSuffix(url, "/") {
		result.URL = result.URL + "/"
	}

	return result
}

func newSendRSA(url, account, password, publicKey string) *Sending {
	result := newSend(url, account, password)
	if len(publicKey) > 0 {
		if !strings.HasPrefix(publicKey, "-----BEGIN PUBLIC KEY-----") {
			publicKey = "-----BEGIN PUBLIC KEY-----\r\n" + publicKey
		}
		if !strings.HasSuffix(publicKey, "-----END PUBLIC KEY-----") {
			publicKey = publicKey + "\r\n" + "-----END PUBLIC KEY-----"
		}
		result.PublicKey = []byte(publicKey)
	}
	return result
}
