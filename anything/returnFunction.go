package anything

import "fmt"

func decorator(innerFunc func(int) int) func(int) int {
	memory := map[int]int{}
	return func(a int) int {
		var ok bool
		var val int
		val = innerFunc(a)
		if _, ok = memory[a]; !ok {
			fmt.Printf("%v is not cached\n", a)
			memory[a] = val
		} else {
			fmt.Printf("%v is cached\n", a)
		}
		return a
	}
}

