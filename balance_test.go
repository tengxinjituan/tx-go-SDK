package apis

import (
	"fmt"
	"testing"
)

// TestBalanceQuery 余额查询测试
func Test_BalanceQuery(t *testing.T) {
	balance := NewBalance("10.10.50.3:8082", "520520", "520520_520520", TestPublicKey)
	model := balance.NewModel("")
	aresult, aerr := balance.Query(model)

	for _, b := range aresult.Balances {
		fmt.Println(b)
	}
	t.Log(aresult, aerr)
	fmt.Println(aresult, aerr)
}

// TestBalanceWatch 余额轮询测试
// 查询三次会出现异常
// panic: test timed out after 30s
// test超过30秒会报panic异常
func Test_BalanceWatch(t *testing.T) {
	balance := NewBalance("10.10.50.3:8082", "520520", "520520_520520", TestPublicKey)
	balance.Watch(10)
}
