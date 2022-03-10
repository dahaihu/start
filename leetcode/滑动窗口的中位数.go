package leetcode

import (
	"fmt"
	"sort"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	mark := make([]int, 0, k)
	for i := 0; i < k; i++ {
		mark = append(mark, nums[i])
	}
	sort.Ints(mark)
	result := []float64{mid(mark)}
	for i := k; i < len(nums); i++ {
		fmt.Printf("original mark is %v\n", mark)
		fmt.Printf("remove ele %d, insert target %d\n", nums[i-k], nums[i])
		insertTarget := sort.Search(k,
			func(j int) bool { return mark[j] > nums[i] })
		fmt.Printf("isnret indx %d\n", insertTarget)
		removeTarget := sort.Search(k,
			func(j int) bool { return mark[j] <= nums[i-k] })
		fmt.Printf("remove indx %d\n", removeTarget)
		if insertTarget > removeTarget {
			for j := removeTarget; j < insertTarget; j++ {
				mark[j] = mark[j+1]
			}
		} else if insertTarget < removeTarget {
			for j := removeTarget; j > insertTarget; j-- {
				mark[j] = mark[j-1]
			}
		}
		fmt.Println(mark, nums[i], nums[i-k], insertTarget, removeTarget)
		mark[insertTarget] = nums[k]
		result = append(result, mid(mark))
	}
	return result
}

func mid(nums []int) float64 {
	res := nums[len(nums)/2]
	if len(nums)%2 == 0 {
		return float64(res)
	}
	return (float64(nums[len(nums)-1]) + float64(res)) / 2
}
