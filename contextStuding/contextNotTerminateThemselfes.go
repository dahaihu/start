package contextStuding

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2020-09-28 15:56
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func inner(ctx context.Context, res chan int) {
	time.Sleep(time.Duration(rand.Int63n(300)) * time.Millisecond)
	select {
	case <- ctx.Done():
		fmt.Println("inner timeout")
	case res <- 100:
		fmt.Println("send to channel")
	}
}

func outer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(150 * time.Millisecond))
	defer cancel()
	res := make(chan int)
	go inner(ctx, res)
	select {
	case <-ctx.Done():
		fmt.Println("out timeout")
	case val := <- res:
		fmt.Println("received value is ", val)
	}
	time.Sleep(time.Millisecond * 100)
}


