package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2020-02-17 08:24
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

type Car struct {
	model string
}

func (c *Car) PrintModel() {
	fmt.Println(c.model)
}


func DeferExp() {
	c := Car{"bm"}
	defer c.PrintModel()
	c.model = "bc"
}
