package leetcode

import (
	"fmt"
	"testing"
)

func TestPTT(t *testing.T) {
	fmt.Println(ptt("ABCDABD"))
	fmt.Println(kmp("BBC ABCDAB ABCDABCDABDE", "ABCDABD"))
}
