package apis

import "github.com/zsy619/sms870/apis/models"

// SMS api接口聚合
type SMS struct {
	sendBatch    *MsgSend
	sendPackage  *MsgPackage
	sendVariable *MsgVariable
	balance      *Balance
}

// NewSMS 初始化api接口
func NewSMS(url, account, password, publicKey string) *SMS {
	result := &SMS{}
	result.sendBatch = NewMsgSend(url, account, password, publicKey)
	result.sendPackage = NewMsgPackage(url, account, password, publicKey)
	result.sendVariable = NewMsgVariable(url, account, password, publicKey)
	result.balance = NewBalance(url, account, password, publicKey)
	return result
}

// NewMsgSendModel 初始化批量发送实体
func (s *SMS) NewMsgSendModel(msg, phone, sendtime, extend, uid string) *models.MsgSend {
	return models.NewMsgSend(s.sendBatch.Account, s.sendBatch.Password, msg, phone, sendtime, extend, uid)
}

// MsgSend 短信发送
func (s *SMS) MsgSend(model *models.MsgSend) (*models.MsgSendResult, error) {
	return s.sendBatch.Send(model)
}

// MsgSendWithRsa RSA加密短信发送
func (s *SMS) MsgSendWithRsa(model *models.MsgSend) (*models.MsgSendResult, error) {
	return s.sendBatch.SendWithRsa(model)
}

// NewMsgPackageModel 初始化短信包发送实体
func (s *SMS) NewMsgPackageModel(msg, sendtime, extend, uid string) *models.MsgPackage {
	return models.NewMsgPackage(s.sendPackage.Account, s.sendPackage.Password, msg, sendtime, extend, uid)
}

// MsgPackage 短信包短信发送
func (s *SMS) MsgPackage(model *models.MsgPackage) (*models.MsgPackageResult, error) {
	return s.sendPackage.Send(model)
}

// MsgPackageWithRsa RSA加密短信包短信发送
func (s *SMS) MsgPackageWithRsa(model *models.MsgPackage) (*models.MsgPackageResult, error) {
	return s.sendPackage.SendRSA(model)
}

// NewMsgVariableModel 初始化变量发送实体
func (s *SMS) NewMsgVariableModel(msg, params, sendtime, extend, uid string) *models.MsgVariable {
	return models.NewMsgVariable(s.sendVariable.Account, s.sendVariable.Password, msg, params, sendtime, extend, uid)
}

// MsgVariable 变量短信发送
func (s *SMS) MsgVariable(model *models.MsgVariable) (*models.MsgVariableResult, error) {
	return s.sendVariable.Send(model)
}

// MsgVariableWithRsa RSA加密变量短信发送
func (s *SMS) MsgVariableWithRsa(model *models.MsgVariable) (*models.MsgVariableResult, error) {
	return s.sendVariable.SendWithRsa(model)
}

// NewBalanceModel 余额查询实体
func (s *SMS) NewBalanceModel(uid string) *models.BalanceModel {
	return models.NewBalanceModel(s.balance.Account, s.balance.Password, uid)
}

// QueryBalance 余额查询
func (s *SMS) QueryBalance(model *models.BalanceModel) (*models.BalanceResult, error) {
	return s.balance.Query(model)
}

// QueryBalanceWithRsa 余额查询
func (s *SMS) QueryBalanceWithRsa(model *models.BalanceModel) (*models.BalanceResult, error) {
	return s.balance.QueryWithRsa(model)
}
