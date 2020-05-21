package playball

import (
	"fmt"
	"math/rand"
	"sync"
)

var Wg sync.WaitGroup

func Play(name string, court chan int) {
	defer Wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}
		n := rand.Int63n(100)
		if n % 13 == 0 {
			close(court)
			fmt.Printf("Player %s lose\n", name)
			return
		}
		fmt.Printf("Player %s hit ball %d\n", name, ball)
		ball++
		court <- ball
	}
}

func test() {

}
