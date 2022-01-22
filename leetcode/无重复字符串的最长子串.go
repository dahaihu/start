package leetcode

func lengthOfLongestSubstring(s string) int {
	var (
		startIdx  int
		maxLength int
	)
	existed := make(map[byte]struct{})
	for idx := 0; idx < len(s); idx++ {
		if _, ok := existed[s[idx]]; ok {
			for s[startIdx] != s[idx] {
				delete(existed, s[startIdx])
				startIdx += 1
			}
			startIdx += 1
		}
		existed[s[idx]] = struct{}{}
		if length := len(existed); length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
