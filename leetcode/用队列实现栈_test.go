package leetcode

import (
	"fmt"
	"testing"
)

func Test_Stack(t *testing.T) {
	stack := MyStack{}
	for _, num := range []int{1, 2, 3, 4} {
		stack.Push(num)
	}
	for !stack.Empty() {
		fmt.Println(stack.Pop())
	}
}
