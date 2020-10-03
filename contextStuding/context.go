package contextStuding

import (
	"context"
	"fmt"
	"time"
)

func Context() {
	gen := func(ctx context.Context) <-chan int {
		res := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case res <- n:
					n++
				}
			}
		}()
		return res
	}
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("ctx type is ", ctx)
	defer cancel()
	for i := range gen(ctx) {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}

func doSomething(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			return
		default:
			fmt.Println("do something ", i)
			i++
		}
	}
}

func nilChannel() <-chan struct{} {
	return nil
}

func GetNilChannel() {
	d := nilChannel()
	fmt.Println("aaa is ", d)
}

func ContextWithCancel() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(parentCtx, "a", "b")
	ctx, cancel := context.WithCancel(valueCtx)
	fmt.Println(ctx == parentCtx)
	go doSomething(ctx)
	time.Sleep(time.Duration(time.Millisecond * 10))
	cancel()
	parentCancel()

}


