package anything

import (
	"fmt"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	fmt.Println(str2bytes("123"))
}

func TestBytes2str(t *testing.T) {
	fmt.Println(bytes2str([]byte{49, 50, 51}))
}