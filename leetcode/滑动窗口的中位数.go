package leetcode

import "sort"

type window []int

func (w *window) remove(ele int) {
	rmPos := sort.Search(len(*w),
		func(idx int) bool { return ele <= (*w)[idx] })
	copy((*w)[rmPos:len(*w)-1], (*w)[rmPos+1:len(*w)])
}

func (w *window) add(ele int) {
	addPos := sort.Search(len(*w)-1,
		func(idx int) bool { return ele <= (*w)[idx] })
	copy((*w)[addPos+1:len(*w)], (*w)[addPos:len(*w)-1])
	(*w)[addPos] = ele
}

func (w *window) mid() float64 {
	length := len(*w)
	if length%2 == 0 {
		return float64((*w)[length/2]+(*w)[(length/2)-1]) / 2
	}
	return float64((*w)[length/2])
}

func medianSlidingWindow(nums []int, k int) []float64 {
	queue := make(window, k)
	copy(queue, nums[:k])
	sort.Ints(queue)
	var result []float64
	for i := k; i < len(nums); i++ {
		result = append(result, queue.mid())
		queue.remove(nums[i-k])
		queue.add(nums[i])
	}
	result = append(result, queue.mid())
	return result
}

type slidingWindow []int



