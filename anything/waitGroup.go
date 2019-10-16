package anything

import (
	"fmt"
	"sync"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-10-11 09:30
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 如果需要给函数传递一个waitgroup，必须传递该值的指针
func workerWait(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func WaitGroupExp() {
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go workerWait(i, &wg)
	}
	wg.Wait()
}
