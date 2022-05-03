package leetcode

import (
	"bytes"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	mark := make(map[string][]string)
	for _, str := range strs {
		to := mapStr(str)
		mark[to] = append(mark[to], str)
	}
	result := [][]string{}
	for _, groups := range mark {
		result = append(result, groups)
	}
	return result
}

var cs = []byte("abcdefghijklmnopqrstuvwxyz")

func mapStr(str string) string {
	count := make([]int, 26)
	for _, s := range str {
		count[s-'a'] += 1
	}
	result := new(bytes.Buffer)
	for idx, cnt := range count {
		if cnt == 0 {
			continue
		}
		result.WriteByte(cs[idx])
		result.WriteString(strconv.Itoa(cnt))
	}
	return result.String()
}
