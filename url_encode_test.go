package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	params := url.Values{}
	params.Add("name", "@Raj eev")
	params.Add("phone", "+919999999999")

	fmt.Println(params.Encode())
}
