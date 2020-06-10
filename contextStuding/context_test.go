package contextStuding

import (
	"context"
	"fmt"
)

// This example demonstrates the use of a cancelable context to prevent a
// goroutine leak. By the end of the example function, the goroutine started
// by gen will return without leaking.
func ExampleWithCancel() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		res := make(chan int)
		var n int = 1
		go func() {
			for {
				select {
				case <- ctx.Done():
					// return so not leaking goroutine
					return
				default:
					res <- n
					n += 1
				}
			}
		}()
		return res
	}
	ctx, cancel := context.WithCancel(context.Background())
	for i := range gen(ctx) {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
	cancel()

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}