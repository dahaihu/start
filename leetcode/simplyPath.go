package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	items := strings.Split(path, "/")
	res := make([]string, 0, 0)
	for _, item := range items {
		if item == "." {
			continue
		} else if item == ".." {
			if len(res) >= 1 {
				res = res[:len(res)-1]
			}
		} else if item != "" {
			res = append(res, item)
		}
	}
	return "/" + strings.Join(res, "/")
}
