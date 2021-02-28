package leetcode

import (
	"fmt"
	"sync"
)

func insertSort(arr []int) {
	for start, end := 1, len(arr); start < end; start++ {
		for i := start; i > 0 && arr[i] < arr[i-1]; i-- {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}
	}
}


func main() {
	c := make(chan string)
	close(c)
	for data := range c {
		fmt.Println(data)
	}
}


func readFromChannel(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range c {
		fmt.Println(data)
	}
}