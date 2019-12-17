package models

// BalanceResult 余额查询结果
type BalanceResult struct {
	Time     string    `json:"time"`     // 响应时间，格式：20180419103048
	Code     string    `json:"code"`     // 状态码：0表示成功
	ErrorMsg string    `json:"errorMsg"` // 状态码说明（成功返回空）
	Balance  string    `json:"balance"`  // 总余额
	Balances []Balance `json:"balances"` // 查询出的余额数组（uid为产品编号, available为可用余额, advance为预支余额）
}

// Balance 余额明细
type Balance struct {
	UId       string `json:"uid"`       // 产品编号
	Available string `json:"available"` // 可用余额
	Advanced  string `json:"advanced"`  // 预支余额
}

// EmptyBalanceResult 返回默认BalanceResult
func EmptyBalanceResult() *BalanceResult {
	return &BalanceResult{}
}

type BalanceModel struct {
	Account  string `json:"account"`  //用户账号
	Password string `json:"password"` //用户密码
	UId      string `json:"uid"`      //产品编号
	Format   string `json:"format"`   //返回格式
}

func NewBalanceModel(account, password, uid string) *BalanceModel {
	return &BalanceModel{Account: account, Password: password, UId: uid, Format: "json"}
}
