package leetcode

/**
* @Author: 胡大海
* @Date: 2020-03-23 08:53
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 感觉堆排序，特别有意思，不知道为啥，就是觉得有意思

// 时间复杂度应该如何取证呢？？？

func adjust(idx int, margin int, array []int) {
	left := 2*idx + 1
	if left >= margin {
		return
	}
	mxChild := left
	if right := left + 1; right < margin && array[right] > array[left] {
		mxChild = right
	}
	if array[idx] >= array[mxChild] {
		return
	}
	array[idx], array[mxChild] = array[mxChild], array[idx]
	adjust(mxChild, margin, array)

}

func heapSort(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		adjust(i, len(array), array)
	}
	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		adjust(0, i, array)
	}
}

func heapfy(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		adjust(i, len(array), array)
	}
}

func selectHeapSort(nums []int) {
	// first step: make nums to big heap
	for i := 1; i < len(nums); i++ {
		cur, num := i, nums[i]
		for cur != 0 {
			parent := (cur - 1) / 2
			if nums[parent] >= num {
				break
			}
			nums[cur] = nums[parent]
			cur = parent
		}
		nums[cur] = num
	}
	for end := len(nums) - 1; end >= 0; end-- {
		nums[end], nums[0] = nums[0], nums[end]
		cur, num := 0, nums[0]
		for cur < end {
			maxChild := 2*cur + 1
			if maxChild >= end {
				break
			}
			if rightChild := maxChild + 1;
				rightChild < end && nums[rightChild] > nums[maxChild] {
				maxChild = rightChild
			}
			if nums[maxChild] <= num {
				break
			}
			cur, nums[cur] = maxChild, nums[maxChild]
		}
		nums[cur] = num
	}
}

type Heap struct {
	nums []int
}

func NewBigHeap(nums []int) *Heap {
	h := &Heap{nums: nums}
	h.heap()
	return h
}

func (h *Heap) heap() {
	for i := 1; i < len(h.nums); i++ {
		h.adjustFromBottomToTop(i)
	}
}

func (h *Heap) adjustFromBottomToTop(startIdx int) {
	num, j := h.nums[startIdx], startIdx
	for j > 0 {
		parent := (j - 1) / 2
		if h.nums[parent] >= num {
			break
		}
		h.nums[j] = h.nums[parent]
		j = parent
		parent = (j - 1) / 2
	}
	h.nums[j] = num
}

func (h *Heap) adjustFromTopToBottom(adjustIdx, endIdx int) {
	for {
		maxChild := adjustIdx*2 + 1
		if maxChild >= endIdx {
			break
		}
		if rightChild := maxChild + 1;
			rightChild < endIdx && h.nums[rightChild] > h.nums[maxChild] {
			maxChild = rightChild
		}
		if h.nums[maxChild] <= h.nums[adjustIdx] {
			break
		}
		h.nums[maxChild], h.nums[adjustIdx] = h.nums[adjustIdx], h.nums[maxChild]
		adjustIdx = maxChild
	}
}

func (h *Heap) Add(array ...int) {
	for _, ele := range array {
		h.nums = append(h.nums, ele)
		h.adjustFromBottomToTop(len(h.nums) - 1)
	}
}

func (h *Heap) Pop() int {
	num := h.nums[0]
	h.nums[0] = h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]
	h.adjustFromTopToBottom(0, len(h.nums)-1)
	return num
}

func (h *Heap) Sort() []int {
	for i := len(h.nums) - 1; i >= 0; i-- {
		h.nums[i], h.nums[0] = h.nums[0], h.nums[i]
		h.adjustFromTopToBottom(0, i)
	}
	return h.nums[:]
}

func (h *Heap) Len() int {
	return len(h.nums)
}