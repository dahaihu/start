package leetcode

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
	fmt.Println(coinChangeUsingDfs([]int{186, 419, 83, 408}, 6249))
	//fmt.Println(coinChangeUsingDfs([]int{1, 3, 4}, 10))
}

func Test_ArrayPointer(t *testing.T) {
	arr := make([]*int, 10)
	for idx, ele := range arr {
		fmt.Println("val is ", ele)
		tmp := idx
		arr[idx] = &tmp
	}
	for _, ele := range arr {
		fmt.Println("val is ", *ele)
	}
}

func Test_FromTopToBottom(t *testing.T) {
	fmt.Println(coinChangeFromTop([]int{2}, 3))
}
