package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-09 09:55
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


func SliceExp() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	s = append(s, "d", "e")
	fmt.Println(s)
	c := make([]string, len(s))
	copy(c, s)
	c[0] = "e"
	fmt.Println(c)
	fmt.Println(s)

	twoD := make([][]int, 3)
	fmt.Println("init slice is ", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
		// 还可以用下面这种方式，作者提供的例子感觉有点容易误导
		// 内部的切片不必重新初始化
		//for j := 0; j < innerLen; j++ {
		//	twoD[i] = append(twoD[i], i + j)
		//}
	}
	fmt.Println(twoD)
}
