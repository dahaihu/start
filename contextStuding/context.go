package contextStuding

import (
	"context"
	"fmt"
)

func Context() {
	gen := func(ctx context.Context) <- chan int {
		res := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <- ctx.Done():
					return
				case res <- n:
					n++
				}
			}
		}()
		return res
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := range gen(ctx) {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}