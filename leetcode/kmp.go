package leetcode

func ptt(s string) []int {
	table := make([]int, 0, len(s))
	for length := 1; length <= len(s); length++ {
		curS := s[:length]
		var wrapLen int
		for ilength := 1; ilength < length; ilength++ {
			if curS[:ilength] == curS[length-ilength:] && ilength > wrapLen {
				wrapLen = ilength
			}
		}
		table = append(table, length-wrapLen)
	}
	return table
}

func kmp(s, p string) bool {
	table := ptt(p)
	i, j := 0, 0
	for j < len(s) {
		if s[j] == p[j-i] {
			j++
		} else {
			if i == j {
				i, j = i+1, j+1
			} else {
				i += table[j-i-1]
			}
		}
		if j-i == len(p) {
			return true
		}
	}
	return j-i == len(p)
}
