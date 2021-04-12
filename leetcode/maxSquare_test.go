package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSquare(t *testing.T) {
	values := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(values))
}
