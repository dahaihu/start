package leetcode

import (
	"fmt"
	"testing"
)

/**
* @Author: 胡大海
* @Date: 2020-11-13 17:28
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func TestMergeSortedArrayList(t *testing.T) {
	arr1 := []int{1, 2, 5, 0, 0, 0}
	arr2 := []int{-1, 2, 4}
	mergeSortedArrayList(arr1, arr2)
	fmt.Println(arr1)
}
