package anything

import (
	"fmt"
	"math"
)

/**
* @Author: 胡大海
* @Date: 2019-10-09 21:42
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

type geometry interface {
	area() float64
	perm() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perm() float64 {
	return 2 * math.Pi * c.radius
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perm() float64 {
	return 2*r.height + 2*r.width
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perm())
}

func InterfaceExp() {
	r := rect{3, 4}
	c := circle{10}
	measure(r)
	measure(c)
}
