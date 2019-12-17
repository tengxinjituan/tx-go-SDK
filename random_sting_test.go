package utils

import (
	"fmt"
	"testing"
)

func Test_RandStringRunes(t *testing.T) {
	key1 := RandStringRunes(16)
	fmt.Println(key1)
}
