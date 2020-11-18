package leetcode

import (
	"bytes"
	"strconv"
)

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//示例 2:
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//说明：
//
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/multiply-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func multiply(num1 string, num2 string) string {
	length1, length2 := len(num1), len(num2)
	mark := make([]int, length1+length2)
	for i := length1 - 1; i >= 0; i-- {
		for j := length2 - 1; j >= 0; j-- {
			val := (num1[i] - '0') * (num2[j] - '0')
			mark[length1-1-i+(length2-1)-j] += int(val)
		}
	}
	carry := 0
	for i := 0; i < len(mark); i++ {
		mark[i] = mark[i] + carry
		carry, mark[i] = mark[i]/10, mark[i]%10
	}
	for len(mark) != 0 && mark[len(mark)-1] == 0 {
		mark = mark[:len(mark)-1]
	}

	if len(mark) == 0 {
		return "0"
	}

	buf := bytes.Buffer{}
	for i := len(mark) - 1; i >= 0; i-- {
		buf.WriteString(strconv.Itoa(mark[i]))
	}
	return buf.String()
}
