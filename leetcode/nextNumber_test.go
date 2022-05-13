package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func Test_nextNumber(t *testing.T) {
	fmt.Println(nextNumber([]int{5, 6, 7}, 78))
}

func Test_binarySearch(t *testing.T) {
	arr := []int{1, 2, 2, 3}
	fmt.Println(sort.Search(len(arr), func(i int) bool { return 2 < arr[i] }))
}
