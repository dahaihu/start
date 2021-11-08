package leetcode

/**
* @Author: 胡大海
* @Date: 2020-04-19 15:49
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"strconv"
	"strings"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var tag string
	if num < 0 {
		tag = "-"
		num = (-1) * num
	}
	mark := make([]string, 0, 3)
	var left int
	for num != 0 {
		num, left = num/7, num%7
		mark = append(mark, strconv.Itoa(left))
	}
	left, right := 0, len(mark)-1
	for left < right {
		mark[left], mark[right] = mark[right], mark[left]
		left++
		right--
	}
	return tag + strings.Join(mark, "")
}
