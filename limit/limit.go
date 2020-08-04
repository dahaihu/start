package limit

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func limitExp() int {
	limiter := rate.NewLimiter(1, 10)
	reserve1 := limiter.ReserveN(time.Now(), 10)
	fmt.Println(reserve1.OK())
	start := time.Now()
	reserve2 := limiter.ReserveN(start, 10)
	reserve3 := limiter.ReserveN(start, 8)
	fmt.Println(reserve2.OK(), reserve3.OK())
	reserve2.CancelAt(start.Add(time.Second))
	fmt.Println(limiter.AllowN(start.Add(time.Second+time.Microsecond*30), 1))
	return 10
}
