package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/groupcache/singleflight"
)

func search() (interface{}, error) {
	fmt.Println("start searching")
	time.Sleep(time.Millisecond * 200)
	return 1000, nil
}

func StudyingSingleFlight() {
	fmt.Printf("package variable a is %d\n", a)
	g := singleflight.Group{}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			res, err := g.Do("multi search", search)
			fmt.Println(res, err)
			wg.Done()
		}()
	}
	wg.Wait()
}
