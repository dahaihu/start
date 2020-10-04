package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2020-02-11 16:05
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
*/

type user struct {
	name string
	age  *int
}

type shower interface {
	show()
}

func sendNotification(n shower) {
	n.show()
}

func (u *user) changeName(name string) {
	u.name = name
}

func (u user) show() {
	fmt.Printf("user name %s, user age %d\n", u.name, u.age)
}

func TestPointer() {
	localUser := new(user)
	localUser.name = "hushichang"
	localUser.show()
	localUser.changeName("liman")
	localUser.show()

	localUser2 := user{name: "hudahai"}
	localUser2.changeName("huhuhu")
	localUser2.show()

	sendNotification(localUser2)
}
