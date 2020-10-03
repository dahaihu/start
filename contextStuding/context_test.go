package contextStuding

import (
	"context"
	"fmt"
	"testing"
)

// This example demonstrates the use of a cancelable context to prevent a
// goroutine leak. By the end of the example function, the goroutine started
// by gen will return without leaking.
func TestWithCancel(t *testing.T) {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	ctx, cancel := context.WithCancel(context.Background())
	gen := func(ctx context.Context) <-chan int {
		res := make(chan int)
		i := 0
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case res <- i + 1:
					i += 1
				}
			}
		}()
		return res
	}
	for i := range gen(ctx) {
		fmt.Println(i)
		if i == 5 {
			cancel()
			break
		}
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func TestChannel(t *testing.T) {
	outer()
}
