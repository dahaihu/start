package anything

import (
	"fmt"
	"time"
)

func autoLoad(d time.Duration) {
	a := 0
	go func() {
		for {
			a += 1
			fmt.Printf("val is %v\n", a)
			time.Sleep(d)
		}
	}()
}
