package leetcode

import (
	"fmt"
	"testing"
)

func TestGasStation(t *testing.T) {
	gas  := []int{2,3,4}
	cost := []int{3,4,3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
