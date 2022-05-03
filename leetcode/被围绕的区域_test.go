package leetcode

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	mark := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}
	solve(mark)
	for i := 0; i < len(mark); i++ {
		fmt.Println(string(mark[i]))
	}
}
