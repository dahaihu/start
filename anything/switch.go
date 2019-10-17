package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-10-09 09:28
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


func SwitchExp() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Print("one\n")
	case 2:
		fmt.Print("two\n")
	default:
		fmt.Print("none of (one, two)\n")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's is weekday")
	}

	// 这个用法很有意思，相当于多个if else一起使用
	// case中的可以使表达式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(int64(1))
	whatAmI("hey")
}

