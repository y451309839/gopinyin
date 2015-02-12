package gopinyin

import (
	"fmt"
	"testing"
)

func Test_ShouPin(t *testing.T) {
	fmt.Printf("转换首拼 Test_ShouPin: Go语言 --> %s\n", ConvertSP("Go语言"))
}

func Test_PinYin(t *testing.T) {
	fmt.Printf("转换拼音 Test_PinYin: Go语言 --> %s\n", ConvertPY("Go语言"))
}
