package apis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zsy619/sms870/apis/models"
	"github.com/zsy619/sms870/utils"
)

// Balance 余额查询
type Balance struct {
	*Sending
}

// NewBalance 初始化余额查询功能
func NewBalance(url, account, password, publicKey string) *Balance {
	api := newSendRSA(url, account, password, publicKey)
	result := &Balance{}
	result.Sending = api
	// 变量短信发送地址
	result.sendURL = result.URL + "msg/balance/json"
	result.sendRSAUrl = result.URL + "msg/balance/json/rsa"
	return result
}

// NewModel 初始化发送实体
func (s *Balance) NewModel(uid string) *models.BalanceModel {
	return models.NewBalanceModel(s.Account, s.Password, uid)
}

// Query 余额查询
func (s *Balance) Query(model *models.BalanceModel) (*models.BalanceResult, error) {
	jsonResult, err := json.Marshal(model)
	if err != nil {
		empty := &models.BalanceResult{}
		return empty, err
	}

	result, err := utils.Post(s.sendURL, string(jsonResult))
	if err != nil {
		empty := &models.BalanceResult{}
		return empty, err
	}
	fmt.Println("query:", result)
	var sendResult models.BalanceResult
	err = json.Unmarshal([]byte(result), &sendResult)
	if err != nil {
		empty := &models.BalanceResult{}
		return empty, err
	}
	return &sendResult, nil
}

// QueryWithRsa 余额查询
func (s *Balance) QueryWithRsa(model *models.BalanceModel) (*models.BalanceResult, error) {
	//对aes的key进行rsa加密
	aesKey := utils.RandStringRunes(16) // aes16位密钥
	key, err := utils.RsaEncrypt([]byte(aesKey), s.PublicKey)
	if err != nil {
		empty := models.EmptyBalanceResult()
		return empty, err
	}
	keyx := base64.StdEncoding.EncodeToString(key)
	rsaSend := &models.RsaSend{
		Account: s.Account,
		Key:     keyx,
	}
	fmt.Println("key:", keyx)

	//对整体内容进行aes加密
	jsonResult, err := json.Marshal(s)
	if err != nil {
		empty := models.EmptyBalanceResult()
		return empty, err
	}
	data, err := utils.AesEncrypt(jsonResult, []byte(aesKey))
	if err != nil {
		empty := models.EmptyBalanceResult()
		return empty, err
	}
	datax := base64.StdEncoding.EncodeToString(data)
	fmt.Println("data:", datax)
	rsaSend.Data = datax

	// 发送
	jsonResultx, errx := json.Marshal(rsaSend) //转换为json格式
	if errx != nil {
		empty := models.EmptyBalanceResult()
		return empty, errx
	}
	fmt.Println("address:", s.sendRSAUrl)
	result, err := utils.Post(s.sendRSAUrl, string(jsonResultx))
	if err != nil {
		empty := models.EmptyBalanceResult()
		return empty, err
	}
	var sendResult models.BalanceResult
	err = json.Unmarshal([]byte(result), &sendResult)
	if err != nil {
		empty := models.EmptyBalanceResult()
		return empty, err
	}
	return &sendResult, nil
}

// Watch 定时查询余额
func (s *Balance) Watch(tt int) {
	go func() {
		if tt < 10 {
			tt = 10
		}
		model := s.NewModel("")
		for {
			result, err := s.QueryWithRsa(model)
			if err != nil {
				fmt.Println("error")
				time.Sleep(time.Second * time.Duration(tt))
				continue
			}
			fmt.Println(result)
			time.Sleep(time.Second * time.Duration(tt))
		}
	}()
	select {}
}
