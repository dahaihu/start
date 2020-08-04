package limit

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

func TestFraction(t *testing.T){
	var duration time.Duration = 1.299999999 * 1e9
	fmt.Println(oldTokensFromDuration(duration, 0.7692307692307693))
	fmt.Println(newTokensFromDuration(duration, 0.7692307692307693))
}


func TestBug(t *testing.T) {
	fmt.Println(rate.NewLimiter(rate.Limit(0.1), 1).Allow())
	fmt.Println(rate.NewLimiter(rate.Limit(0.2), 1).Allow())
	fmt.Println(rate.NewLimiter(rate.Limit(0.3), 1).Allow())
	fmt.Println(rate.NewLimiter(rate.Limit(0.4), 1).Allow())
	fmt.Println(rate.NewLimiter(rate.Limit(0.5), 1).Allow())
	fmt.Println(rate.NewLimiter(rate.Limit(0.6), 1).Allow())
	fmt.Println(rate.NewLimiter(rate.Limit(0.7), 1).Allow())
	fmt.Println(rate.NewLimiter(rate.Limit(0.8), 1).Allow())
	fmt.Println(rate.NewLimiter(rate.Limit(0.9), 1).Allow())

	fmt.Println(rate.NewLimiter(rate.Limit(0.7692307692307693), 1).Allow())
}