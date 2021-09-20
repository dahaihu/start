package anything

import (
	"fmt"
	"testing"
)

func findWildcard(path string) (wildcard string, i int, valid bool) {
	// Find start
	for start, c := range []byte(path) {
		// A wildcard starts with ':' (param) or '*' (catch-all)
		if c != ':' && c != '*' {
			continue
		}

		// Find end and check for invalid characters
		valid = true
		for end, c := range []byte(path[start+1:]) {
			switch c {
			// 找到的 wildcard 是 /:name/age 里面的 :name
			case '/':
				return path[start : start+1+end], start, valid
			case ':', '*':
				valid = false
			}
		}
		// 找到的 wildcard 是 /:name 里面的 :name
		return path[start:], start, valid
	}
	return "", -1, false
}

func TestFindWildcard(t *testing.T) {
	path := "/internal/*encoding"
	wildcard, idx, valid := findWildcard(path)
	fmt.Printf("wildcard is %v, idx value is %v, valid is %v\n", wildcard, path[:idx], valid)
}
