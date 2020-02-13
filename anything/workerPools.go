package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2020-02-12 21:32
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// oneWorker 既然相当于工人， 那么他需要什么呢？输入，输出，以及确认它是哪一个工人
func oneWorker(id int, jobs <-chan int, results chan<- int) {
	for  {
		job, ok := <- jobs
		if !ok {
			break
		}
		fmt.Printf("worker %d start job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %d end job %d\n", id, job)
		results <- 2 * job
	}
}

func WorkerPoolExp() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for i := 1; i <= 3; i++ {
		go oneWorker(i, jobs, results)
	}

	for i := 1; i<= 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i<= 5; i++ {
		<-results
	}
}
