package leetcode

import (
	"bytes"
	"strconv"
)

/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func GetPermutation(n int, k int) string {
	res := bytes.Buffer{}
	mark := make([]int, n)
	items := make([]int, n)
	for i:=0; i<n; i++ {
		if i==0 {
			mark[i] = 1
		} else {
			mark[i] = i * mark[i-1]
		}
		items[i] = i+1
	}
	var idx int
	for {
		idx, k = k/mark[len(mark)-1], k%mark[len(mark)-1]

		if k == 0 {
			res.WriteString(strconv.Itoa(items[idx-1]))
			items = pop(items, idx-1)
			for i := len(items)-1; i>=0; i-- {
				res.WriteString(strconv.Itoa(items[i]))
			}
			break
		}
		res.WriteString(strconv.Itoa(items[idx]))
		items = pop(items, idx)
		mark = mark[:len(mark)-1]
	}
	return res.String()
}

func pop(items []int, idx int) []int {
	res := make([]int, len(items)-1)
	copy(res[:idx], items[:idx])
	copy(res[idx:], items[idx+1:])
	return res
}
