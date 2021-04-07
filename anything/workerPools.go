package anything

import (
	"fmt"
	"sync"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2020-02-12 21:32
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// oneWorker 既然相当于工人， 那么他需要什么呢？输入，输出，以及确认它是哪一个工人
func oneWorker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		job, ok := <- jobs
		if !ok {
			return
		}
		fmt.Printf("worker %d start doing job %d\n", id, job)
		time.Sleep(time.Microsecond * 100)
		fmt.Printf("worker %d has done the  job %d\n", id, job)
	}
}

func WorkerPoolExp() {
	const numJobs = 100
	const workers = 3
	jobs := make(chan int, 3)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func () {
		defer wg.Done()
		for i := 0; i < numJobs; i++ {
			fmt.Printf("sending job %d\n", i)
			jobs <- i
		}
		close(jobs)
	}()
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go oneWorker(i, jobs, wg)
	}
	wg.Wait()
	fmt.Println("wait done ?????")
}
