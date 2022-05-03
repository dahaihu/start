package leetcode

import (
	"fmt"
	"testing"
)

func TestConvertToBase7(t *testing.T) {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
}

func TestSlice(t *testing.T) {
	a := []int{}
	fmt.Println(len(a), cap(a))
	b := append(a, 1)
	fmt.Println(len(a), cap(a), len(b), cap(b))
	//a = append(a, 1)
	//fmt.Println(len(a), cap(a))
	//a = append(a, 1)
	//fmt.Println(len(a), cap(a))
}