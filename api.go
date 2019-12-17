package apis

import (
	"github.com/zsy619/sms870/apis/models"
)

const (
	//TestPublicKey rsa测试公钥
	TestPublicKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDYI66GpjudXNiOdHb4yEfw6m1d IKUppz4Fc56iTT3ABXtrkpfqYgs4/kJZEfgGGcuDnKVmBEmQ7PcCEe6hJv3Vh+wg 7U7Uv5FXO4t1sUoWMaqR8a/p8HLrtU3pGSGIrVniCbXQZi5kAVJaJ8Heae8fMrtG X9+Zp1QkujZWDg3A0QIDAQAB`
)

// API 接口定义
type API interface {
	NewSendModel(msg, sendtime, extend, uid string) *models.Send
	Send(model *models.Send) (*models.SendResult, error)
	SendRSA(model *models.Send) (*models.SendResult, error)
}
