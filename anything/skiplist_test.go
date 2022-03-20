package anything

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	s := Constructor()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println(&s)
	fmt.Println(s.Search(0))
	s.Add(4)
	fmt.Println(&s)
	fmt.Println(s.Search(1))
	fmt.Println(s.Erase(0))
	fmt.Println(&s)
	fmt.Println(s.Erase(1))
	fmt.Println(&s)
	fmt.Println(s.Search(1))
	//fmt.Println("************inserted**************")
	//for i := 0; i < len(s.Head.Next); i++ {
	//	s.PrintLevel(i)
	//}
	//for i := 0; i < nums; i++ {
	//	fmt.Println(s.Delete(i))
	//}
	//fmt.Println("************deleted**************")
	//for i := 0; i < len(s.Head.Next); i++ {
	//	s.PrintLevel(i)
	//}
}
