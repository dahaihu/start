package anything

import "fmt"

func gotoStudying() {
	var a = 10
LOOP:
	if a < 20 {
		a += 1
		goto LOOP
	}
	fmt.Printf("final a value is %v\n", a)
}