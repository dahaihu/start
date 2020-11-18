package leetcode

import (
	"testing"
)

func TestAddTwoList(t *testing.T) {
	a := initiateList([]int{2, 4, 9})
	b := initiateList([]int{5, 6, 4, 9})
	res := addTwoNumbers(a, b)
	traverseList(res)
}

