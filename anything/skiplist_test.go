package anything

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSkipList(t *testing.T) {
	s := Constructor()
	for i := 1; i <= 20; i++ {
		s.Add(i)
		if rand.Float64() < 0.5 {
			s.Add(i)
		}
	}
	fmt.Println(&s)
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
