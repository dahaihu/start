package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-15 13:37
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func InitiateExp() {
	// 不赋初值的边纳凉，最好通过var声明
	var a []string
	a = append(a, "张三")
	fmt.Println(len(a), cap(a))
	b := []string{}
	b = append(b, "李四")
	fmt.Println(len(b), cap(b))
}
