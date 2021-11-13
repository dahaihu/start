package leetcode

import (
	"strconv"
	"strings"
)

/**
* @Author: 胡大海
* @Date: 2020-04-19 15:49
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

/*
给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。

 

示例 1:

输入: num = 100
输出: "202"
示例 2:

输入: num = -7
输出: "-10"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/base-7
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var (
		tag  string
		left int
	)
	if num < 0 {
		tag = "-"
		num = (-1) * num
	}
	mark := make([]string, 0, 4)
	for num != 0 {
		left, num = num%7, num/7
		mark = append(mark, strconv.Itoa(left))
	}
	leftIdx, rightIdx := 0, len(mark)-1
	for leftIdx < rightIdx {
		mark[leftIdx], mark[rightIdx] = mark[rightIdx], mark[leftIdx]
		leftIdx++
		rightIdx--
	}
	return tag + strings.Join(mark, "")
}
