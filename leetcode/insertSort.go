package leetcode

import (
	"fmt"
	"sync"
)

func insertSort(arr []int) {
	for dummy := 1; dummy < len(arr); dummy++ {
		for mark := dummy-1; mark >= 0 && arr[mark] > arr[mark+1]; mark-- {
			arr[mark], arr[mark+1] = arr[mark+1], arr[mark]
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