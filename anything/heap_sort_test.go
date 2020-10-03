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

func RandomSequence(min, max int) []int {
	//计算序列的长度
	lenghth := max - min + 1

	//初始化一个长度为lenghth的原始切片，初始值从min到max
	initArr := make([]int, lenghth)
	for i := 0; i < lenghth; i++ {
		initArr[i] = i + min
	}

	//初始化一个长度为lenghth的目标切片
	rtnArr := make([]int, lenghth)

	//初始化随机种子
	rand.Seed(time.Now().Unix())

	//生成目标序列
	for i := 0; i < lenghth; i++ {
		//生成一个随机序号
		index := rand.Intn(lenghth - i)

		//将原始切片中序号index对应的值赋给目标切片
		rtnArr[i] = initArr[index]

		//替换掉原始切片中使用过的下标index对应的值
		initArr[index] = initArr[lenghth-i-1]
	}

	return rtnArr
}


func TestSequence(t *testing.T) {
	mark := RandomSequence(1, 10)
	fmt.Println("original mark is ", mark)
	heapSort(mark)
	fmt.Println("sorted mark is ", mark)

}
