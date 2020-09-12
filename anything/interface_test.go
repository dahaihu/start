package anything

import "testing"

func TestInterface(t *testing.T) {
	r := rect{3, 4}
	c := circle{10}
	s := square{100}
	measure(r)
	measure(c)
	measure(s)
}