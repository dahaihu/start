package leetcode

import (
	"math/rand"
	"testing"
	"time"
)

func randomSequence(min, max int) []int {
	//计算序列的长度
	length := max - min + 1

	//初始化一个长度为length的原始切片，初始值从min到max
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i + min
	}

	//初始化随机种子
	rand.Seed(time.Now().Unix())

	for i := length - 1; i >= 0; i-- {
		idx := rand.Intn(i + 1)
		arr[idx], arr[i] = arr[i], arr[idx]
	}

	return arr
}

func invalidSortedArray(nums []int) bool {
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < pre {
			return true
		}
		pre = nums[i]
	}
	return false
}

func TestQuickSort(t *testing.T) {
	for i := 0; i <= 1000; i++ {
		nums := randomSequence(1, 10)
		nums[5] = 5
		quickMain(nums)
		if invalidSortedArray(nums) {
			t.Fatalf("invalid sorted array %v", nums)
		}
	}
}


func TestSplit(t *testing.T) {
	//nums := randomSequence(1, 10)
	//nums[5] = 5
	//right, left := splitArray(nums, 5)
	//fmt.Println(right, left, nums)

}
