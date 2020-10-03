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
	time.Sleep(time.Duration(rand.Int63n(600)) * time.Millisecond)
	select {
	case <- ctx.Done():
		fmt.Println("inner timeout")
	case res <- 100:
		fmt.Println("inner done")
	}
}

func outer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(300)*time.Millisecond)
	defer cancel()
	res := make(chan int)
	go inner(ctx, res)
	select {
	case <-ctx.Done():
		fmt.Println("outer done")
	case val := <-res:
		fmt.Println("outer finished, result is ", val)
	}
	time.Sleep(time.Duration(300) * time.Millisecond)
}
