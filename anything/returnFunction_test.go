package anything

import "testing"

func TestReturnFunc(t *testing.T) {
	original := decorator(func(a int) int {return a})
	original(10)
	original(10)
	last := decorator(func(a int) int {return a})
	last(10)
	last(10)
}
