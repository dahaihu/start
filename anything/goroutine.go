package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-10-11 07:59
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */
func f(from string) {
	for i:=0; i<3; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(from, ":", i)
	}
}

func GoroutineExp() {
	f("direct")
	go f("synchronious")
	go func() {
		fmt.Println("hehe")
	}()
	time.Sleep(time.Second)
	fmt.Println("done")
}
