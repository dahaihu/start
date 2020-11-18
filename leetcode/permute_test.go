package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	values := []int{1, 2, 3}
	fmt.Println(permute(values))
}

func TestSort(t *testing.T) {
	values := []int{3, 1, 4, 1, 10}
	sort.Ints(values)
	fmt.Println(values)
}

func ExampleGetPermutation() {
	a := [3]int{1, 2, 3}
	b := a[:2:3]
	b = append(b, 4)
	//fmt.Println(a)
	fmt.Println(GetPermutation(4, 9))
	// Output: 2314
}