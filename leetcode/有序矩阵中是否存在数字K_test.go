package leetcode

import (
	"fmt"
	"sync/atomic"
	"testing"
	"unsafe"
)

/**
* @Author: 胡大海
* @Date: 2020-11-18 14:57
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

type Test struct {
	age int
}


func TestExistValInMultiArray(t *testing.T) {
	test := Test{age: 10}
	fmt.Println((&test).age)
	fmt.Println(&test.age)
	a := 10
	b := 20
	addr := unsafe.Pointer(&a)
	//fmt.Println(&a == &b)
	bddr := unsafe.Pointer(&b)
	fmt.Println("changed ", atomic.CompareAndSwapPointer(&addr, addr, bddr))
	fmt.Println(*(*int)(addr))
	fmt.Println(existValInMultiArray([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, -1))
}
