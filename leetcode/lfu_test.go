package leetcode
//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestLFUCASE(t *testing.T) {
//	lfu := LFUConstructor(3)
//	for i := 0; i <= 10; i++ {
//		lfu.Put(3, 3)
//	}
//	fmt.Println(lfu)
//}
//
//func TestLFU(t *testing.T) {
//	lfu := LFUConstructor(10)
//	lfu.Put(10, 13)
//	lfu.Put(3, 17)
//	lfu.Put(6, 11)
//	lfu.Put(10, 5)
//	lfu.Put(9, 10)
//	fmt.Println(lfu.Get(13) == -1)
//	lfu.Put(2, 19)
//	fmt.Println(lfu.Get(2) == 19)
//	fmt.Println(lfu.Get(3) == 17)
//	lfu.Put(5, 25)
//	fmt.Println(lfu.Get(8) == -1)
//	lfu.Put(9, 22)
//	lfu.Put(5, 5)
//	lfu.Put(1, 30)
//	fmt.Println(lfu.Get(11) == -1)
//	lfu.Put(9, 12)
//	//fmt.Println(lfu.Get(7) == -1)
//	//fmt.Println(lfu.Get(5) == 5)
//	//fmt.Println(lfu.Get(8) == -1)
//	//fmt.Println(lfu.Get(9) == 12)
//	//lfu.Put(4, 30)
//	//lfu.Put(9, 3)
//	//fmt.Println(lfu.Get(9) == 3)
//	//fmt.Println(lfu.Get(10) == 5)
//	//fmt.Println(lfu.Get(10) == 5)
//	//lfu.Put(6, 14)
//	//lfu.Put(3, 1)
//	//fmt.Println(lfu.Get(3) == 1)
//	//lfu.Put(10, 11)
//	//fmt.Println(lfu.Get(8) == -1)
//	//lfu.Put(2, 14)
//	//fmt.Println(lfu.Get(1) == 30)
//	//fmt.Println(lfu.Get(5) == 5)
//	//fmt.Println(lfu.Get(4) == 30)
//	//lfu.Put(11, 4)
//	//lfu.Put(12, 24)
//	//lfu.Put(5, 18)
//	//fmt.Println(lfu.Get(13) == -1)
//	//lfu.Put(7, 23)
//	//fmt.Println(lfu.Get(8) == -1)
//	//fmt.Println(lfu.Get(12) == 24)
//	//lfu.Put(3, 27)
//	//lfu.Put(2, 12)
//	//fmt.Println(lfu.Get(5) == 18)
//	//lfu.Put(2, 9)
//	//lfu.Put(13, 4)
//	//lfu.Put(8, 18)
//	//lfu.Put(1, 7)
//	//fmt.Println(lfu.Get(6) == 14)
//	//lfu.Put(9, 29)
//	//lfu.Put(8, 21)
//	//fmt.Println(lfu.Get(5) == 18)
//	//lfu.Put(6, 30)
//	//lfu.Put(1, 12)
//	//fmt.Println(lfu.Get(10) == 11)
//	//lfu.Put(4, 15)
//	//lfu.Put(7, 22)
//	//lfu.Put(11, 26)
//	//lfu.Put(8, 17)
//	//lfu.Put(9, 29)
//	//fmt.Println(lfu.Get(5) == 18)
//	//lfu.Put(3, 4)
//	//lfu.Put(11, 30)
//	//fmt.Println(lfu.Get(12) == -1)
//	//lfu.Put(4, 29)
//	//fmt.Println(lfu.Get(3) == 4)
//	//fmt.Println(lfu.Get(9) == 29)
//	//fmt.Println(lfu.Get(6) == 30)
//	//lfu.Put(3, 4)
//	//fmt.Println(lfu.Get(1) == 12)
//	//fmt.Println(lfu.Get(10) == 11)
//	//lfu.Put(3, 29)
//	//lfu.Put(10, 28)
//	//lfu.Put(1, 20)
//	//lfu.Put(11, 13)
//	//fmt.Println(lfu.Get(3) == 29)
//	//lfu.Put(3, 12)
//	//lfu.Put(3, 8)
//	//lfu.Put(10, 9)
//	//lfu.Put(3, 26)
//	//fmt.Println(lfu.Get(8) == 17)
//	//fmt.Println(lfu.Get(7) == -1)
//	//fmt.Println(lfu.Get(5) == 18)
//	//lfu.Put(13, 17)
//	//lfu.Put(2, 27)
//	//lfu.Put(11, 15)
//	//fmt.Println(lfu)
//	//fmt.Println(lfu.items)
//	//fmt.Println(lfu.Get(12) == -1)
//	//lfu.Put(9, 19)
//	//lfu.Put(2, 15)
//	//lfu.Put(3, 16)
//	//fmt.Println(lfu.Get(1) == 20)
//	//lfu.Put(12, 17)
//	//lfu.Put(9, 1)
//	//lfu.Put(6, 19)
//	//fmt.Println(lfu.Get(4) == 29)
//	//fmt.Println(lfu.Get(5) == 18)
//	//fmt.Println(lfu.Get(5) == 18)
//	//lfu.Put(8, 1)
//	//lfu.Put(11, 7)
//	//lfu.Put(5, 2)
//	//lfu.Put(9, 28)
//	//fmt.Println(lfu.Get(1) == 20)
//	//lfu.Put(2, 2)
//	//lfu.Put(7, 4)
//	//lfu.Put(4, 22)
//	//lfu.Put(7, 24)
//	//lfu.Put(9, 26)
//	//lfu.Put(13, 28)
//	//lfu.Put(11, 26)
//	fmt.Println(lfu)
//}
