package anything

/**
* @Author: 胡大海
* @Date: 2020-10-03 15:29
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"fmt"
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

	for i:=length-1; i >= 0; i-- {
		idx := rand.Intn(i+1)
		arr[idx], arr[i] = arr[i], arr[idx]
	}

	return arr
}


func TestSequence(t *testing.T) {
	mark := randomSequence(1, 10)
	fmt.Println("original mark is ", mark)
	heapSort(mark)
	fmt.Println("sorted mark is ", mark)

}
