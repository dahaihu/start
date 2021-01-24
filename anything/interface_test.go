package anything

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	var r geometry = rect{3, 4}
	var c geometry = circle{10}
	fmt.Println("r == c : ", r == c)
	fmt.Println(1e6 == 1000000)
}