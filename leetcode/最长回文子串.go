package leetcode

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := s[0:1]
	mark := make([][]bool, len(s))
	for i:=0; i<len(s); i++{
		mark[i] = make([]bool, len(s))
		mark[0][i] = true

	}
	for i:=0; i<len(s)-1; i++ {
		mark[1][i] = s[i] == s[i+1]
		if mark[1][i] && len(res) == 1 {
			res = s[i:i+2]
		}
	}
	fmt.Println("mark is ", mark)
	for length := 2; length < len(s); length ++ {
		for start := 0; start <= len(s)-length-1; start++ {
			mark[length][start] = mark[length-2][start+1] && s[start] == s[start+length]
			if mark[length][start] && length+1 > len(res) {
				res = s[start: start+length+1]
			}
		}
	}
	fmt.Println("mark is ", mark)
	return res
}
