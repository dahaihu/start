package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-09 21:55
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 忘记学过的语言，这些东西可能带给你些不好的东西
// go by example，真的是非常好的文章，用于学习和回顾知识
// 另外，作者还会拓展一些知识点，提供一些比较好的博客

func sum(nums ...int) {
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func VariadicFunctionExp() {
	sum(1, 2, 3)
	nums := []int{1, 2, 3}
	sum(nums...)
}
