package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-24 20:50
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


type People struct {}


func (p *People) ShowA() {
	fmt.Println("ShowA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("ShowB")
}


type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Emebeding() {
	t := Teacher{}
	t.ShowA()
}
