package leetcode

import (
	"fmt"
	"testing"
)

import (
	"math/rand"
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

func TestQuickSort(t *testing.T) {
	nums := randomSequence(1, 10)
	fmt.Println("nums is ", nums)
	nums[5] = 5
	quickMain(nums)
}
