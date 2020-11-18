package leetcode

import "strings"

func generateParenthesis(n int) []string {
	res := make([][]string, 0)
	var dep func(map[string]int, []string)
	dep = func(leftMap map[string]int, tmpRes []string) {
		if leftMap["("] == 0 {
			for i := 0; i < leftMap[")"]; i++ {
				tmpRes = append(tmpRes, ")")
			}
			res = append(res, tmpRes)
			return
		}
		if leftMap["("] == leftMap[")"] {
			tmp := make([]string, len(tmpRes)+1)
			copy(tmp, tmpRes)
			tmp[len(tmp)-1] = "("
			dep(map[string]int{"(": leftMap["("] - 1, ")": leftMap[")"]}, tmp)
			return
		}
		for _, ele := range [2]string{"(", ")"} {
			tmp := make([]string, len(tmpRes)+1)
			copy(tmp, tmpRes)
			tmp[len(tmp)-1] = ele
			if ele == "(" {
				dep(map[string]int{"(": leftMap["("] - 1, ")": leftMap[")"]}, tmp)
			} else {
				dep(map[string]int{"(": leftMap["("], ")": leftMap[")"] - 1}, tmp)
			}
		}
	}
	dep(map[string]int{"(": n, ")": n}, make([]string, 0))
	finalRes := make([]string, len(res))
	for idx, elements := range res {
		finalRes[idx] = strings.Join(elements, "")
	}
	return finalRes
}
