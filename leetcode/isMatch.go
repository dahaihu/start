package leetcode

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-11-10 12:15
* A programmer who subconsciously views himself as an artist will enpIdxoy what he does and will do it better ​
 */

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func IsMatch(s string, p string) bool {
	mark := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		mark[i] = make([]bool, len(s)+1)
	}
	mark[0][0] = true

	for i := 1; i < len(p)+1; i++ {
		if p[i-1] == '*' {
			mark[i][0] = mark[i-2][0]
		}
	}

	for pPos := 1; pPos <= len(p); pPos++ {
		for sPos := 1; sPos <= len(s); sPos++ {
			if s[sPos-1] == p[pPos-1] || p[pPos-1] == '.' {
				mark[pPos][sPos] = mark[pPos-1][sPos-1]
			} else if p[pPos-1] == '*' {
				mark[pPos][sPos] = mark[pPos-2][sPos]
				if p[pPos-2] == s[sPos-1] || p[pPos-2] == '.' {
					mark[pPos][sPos] = mark[pPos][sPos] || mark[pPos][sPos-1]
				}
			}
		}
	}
	return mark[len(p)][len(s)]
}

func IsMatch2(s string, p string) bool {
	mark := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		mark[i] = make([]bool, len(s)+1)
	}
	mark[0][0] = true

	for i := 1; i < len(p)+1; i++ {
		if p[i-1] == '*' {
			mark[i][0] = mark[i-1][0]
		}
	}

	for pPos := 1; pPos <= len(p); pPos++ {
		for sPos := 1; sPos <= len(s); sPos++ {
			if s[sPos-1] == p[pPos-1] || p[pPos-1] == '?' {
				mark[pPos][sPos] = mark[pPos-1][sPos-1]
			} else if p[pPos-1] == '*' {
				mark[pPos][sPos] = mark[pPos-1][sPos] ||
					mark[pPos][sPos-1]
			}
		}
	}
	fmt.Println(mark)
	return mark[len(p)][len(s)]
}
