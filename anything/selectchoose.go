package anything

import (
	"fmt"
	"runtime"
)

/**
* @Author: 胡大海
* @Date: 2019-10-24 21:19
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


func SelectChooseOne() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	intChan <- 1
	stringChan <- "zhangsan"
	// select 在都可以有值的时候，会随机选择一个
	// 而不是从第一个case的赋值语句开始判断
	// 这个是值得思考的一个地方，不要有先入为主的思想
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		fmt.Println(value)
	}
}
