package leetcode

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
	//fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(change([]int{1, 2, 5}, 0))
	//fmt.Println(coinChange([]int{2}, 3))
	//fmt.Println(coinChange([]int{1, 2, 5}, 10))
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
	fmt.Println(coinChange([]int{2}, 3))
}
