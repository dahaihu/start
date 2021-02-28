package leetcode

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	a := []int{1, 2, 4, 3, -1}
	insertSort(a)
	//sort.Slice(a, func(i int, j int) bool { return a[i] < a[j] })
	fmt.Println(a)
}

func TestAain(t *testing.T) {
	var c chan string

	select {
	case c <- "123":
		fmt.Println("send value to channel c")
	}
	fmt.Println("final")
}
