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
	items := make([]int, n)
	mark := make([]int, n)
	mark[0] = 1
	items[0] = 1
	for i := 1; i < n; i++ {
		mark[i] = mark[i-1] * i
		items[i] = i + 1
	}
	var idx int
	for true {
		idx, k = k / mark[len(items) - 1], k % mark[len(items) - 1]
		if k == 0 {
			res.WriteString(strconv.Itoa(items[idx - 1]))
			items = popKItem(items, idx-1)
			for i := len(items) - 1; i >=0; i-- {
				res.WriteString(strconv.Itoa(items[i]))
			}
			break
		}
		res.WriteString(strconv.Itoa(items[idx]))
		items = popKItem(items, idx)
	}
	return res.String()
}


func popKItem(items []int, idx int) []int {
	res := make([]int, len(items) - 1)
	copy(res[:idx], items[:idx])
	copy(res[idx:], items[idx+1:])
	return res
}
