package leetcode

/**
* @Author: 胡大海
* @Date: 2019-10-29 21:28
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"fmt"
	"strings"
)

// bad case 做一个博客，介绍一下相关的情况, 不用介绍了，傻逼操作而已
// 输入是这个东西 	ret := Subsets([]int{9,0,3,5,7})

func Subsets(nums []int) [][]int {
	ret := [][]int{{}}
	for _, num := range nums {
		// // 这个copy并没有什么卵用
		//tmp := make([][]int, len(ret))
		//copy(tmp, ret)
		//fmt.Printf("num %d, tmp'length is %d, tmp is %v\n", num, len(tmp), tmp)
		for _, tmpElems := range ret {
			fmt.Println(strings.Repeat("a", 100))
			fmt.Printf("tmpElems's len is %d, cap is %d\n", len(tmpElems), cap(tmpElems))
			fmt.Printf("%d before tmpElems is %v\n", num, tmpElems)
			tmpElems = append(tmpElems, num)
			ret = append(ret, tmpElems)
			fmt.Printf("%d after tmpElems is %v\n", num, tmpElems)
			fmt.Printf("ret is %v\n", ret)
		}
	}
	return ret
}

// 区别是共享一个底层的数组
// a := [4]int{0, 1, 2, 3, 4}
// s := a[1:3]
// s1 := s[1:3]
// s2 := s[1:3]
// s1 = append(s1, 1)
// s2 = append(s2, 1)
// 造成了对共享的数据的修改

func RightSubsets(nums []int) [][]int {
	ret := [][]int{{}}
	for _, num := range nums {
		tmpRet := make([][]int, len(ret))
		copy(tmpRet, ret)
		for _, subset := range tmpRet {
			// 这个地方不进行创建新的，而是在原来的subset进行append的话，可能会对原来的subset进行了修改
			// 因为前面部分的subset底层的array肯能是相同的
			newSubset := make([]int, len(subset)+1)
			copy(newSubset, subset)
			newSubset[len(subset)] = num
			ret = append(ret, newSubset)
		}
	}
	return ret
}
