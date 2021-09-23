package leetcode

func maxLengthPalindrome(s string) string {
	var palindrome string
	if len(s) < 1 {
		return palindrome
	}
	palindrome = s[0:1]
	mark := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		mark[i] = make([]int, len(s))
		mark[i][0] = 1
		if i+1 < len(s) && s[i] == s[i+1] {
			mark[i][1] = 2
			if len(palindrome) == 1 {
				palindrome = s[i : i+2]
			}
		}
	}
	for length := 2; length < len(s); length++ {
		for start := 0; start < len(s)-length; start++ {
			if end := start + length;
				s[start] == s[end] && mark[start+1][length-2] > 0 {
				mark[start][length] = length
				if length+1 > len(palindrome) {
					palindrome = s[start : start+length+1]
				}
			}
		}
	}
	return palindrome
}
