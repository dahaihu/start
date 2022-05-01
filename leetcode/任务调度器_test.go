package leetcode

import (
	"fmt"
	"testing"
)

func Test_leastIntervalBetter(t *testing.T) {
	fmt.Println(leastInterval([]byte{
		'A', 'A', 'A', 'A', 'A', 'A',
		'B', 'B', 'B', 'B', 'B',
		'C', 'C', 'C', 'C',
		'D', 'D', 'D',
		'E', 'E', 'E',
	}, 3))
}