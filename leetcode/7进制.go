package leetcode

/**
* @Author: 胡大海
* @Date: 2020-04-19 15:49
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"bytes"
	"strconv"
)
func ConvertToBase7(num int) string {
	mark := make([]string, 0)
	tag := ""
	if num == 0 {
		return "0"
	} else if num < 0 {
		num = -num
		tag = "-"
	}
	var remaining int
	for ;num != 0; {
		remaining, num = num % 7, num / 7
		mark = append(mark, strconv.Itoa(remaining))
	}
	buffer := bytes.Buffer{}
	buffer.WriteString(tag)
	for i := len(mark) - 1; i >= 0; i-- {
		buffer.WriteString(mark[i])
	}
	return buffer.String()
}
