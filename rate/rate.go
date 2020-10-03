package rate

/**
* @Author: 胡大海
* @Date: 2020-07-25 15:48
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func rateStuding() {
	limiter := rate.NewLimiter(rate.Limit(10), 1000)
	reservation := limiter.ReserveN(time.Now(), 1000)
	fmt.Println(reservation.OK())
	reservation = limiter.ReserveN(time.Now(), 1000)
	fmt.Println(reservation.OK())
	reservation.CancelAt(time.Now().Add(time.Second))
	//fmt.Println(limiter.AllowN(time.Now().Add(time.Second + time.Millisecond * 30), 1500))
}
