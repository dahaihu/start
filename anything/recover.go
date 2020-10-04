package anything

import (
	"fmt"
)

/**
* @Author: 胡大海
* @Date: 2020-02-12 08:31
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func TestRecover() {
	defer func() {
		// 不管是嵌套的多少层，都可以在此进行恢复
		// 也就是说，可以在应用的根部进行如此的调用
		if err := recover(); err != nil {
			fmt.Println("err is ", err)
		}
	}()
	fmt.Println("first")
	customePanic()
	fmt.Println("second")
}

func customePanic() {
	innerPanic()
}

func innerPanic() {
	panic("sb")
}

// 这种东西，感觉纯粹就是一个坑，所以可能成了一个面试题
func cautiousDefer() (i int) {
	// return number will be 2
	defer func() {
		i++
	}()
	return 1
}
