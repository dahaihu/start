package leetcode

import (
	"fmt"
	"testing"
)

func ExampleQueens() {
	fmt.Println(solveNQueens(4))
	// Output: [[1 3 0 2] [2 0 3 1]]
}

func Test_String(t *testing.T) {
	for _, ele := range "zhangsan" {
		fmt.Println("ele is ", ele)
	}
}
