package anything

import "fmt"

func Fallthrough(val bool) {
	// if val is true, then three print will all be executed
	switch val {
	case false:
		fmt.Println("The integer was <= 4")
		fallthrough
	case true:
		fmt.Println("The integer was <= 5")
		fallthrough
	default:
		fmt.Println("default case")
	}
}