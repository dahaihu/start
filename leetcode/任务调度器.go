package leetcode

import "fmt"

func leastInterval(tasks []byte, n int) int {
	taskCnt := make(map[byte]int)
	for _, task := range tasks {
		taskCnt[task] += 1
	}
	cnts := make([]int, 0, len(taskCnt))
	for _, cnt := range taskCnt {
		cnts = append(cnts, cnt)
	}
	result, heap := 0, NewBigHeap(cnts)
	for {
		var tmp []int
		for i := 0; i <= n; i++ {
			if heap.Len() > 0 {
				if ele := heap.Pop(); ele > 1 {
					tmp = append(tmp, ele-1)
				}
			} else if len(tmp) == 0 {
				// heap.Len() == 0 && len(tmp) == 0
				return result
			}
			result += 1
		}
		heap.Add(tmp...)
	}
}

func leastIntervalBetter(tasks []byte, n int) int {
	cnts := make([]int, 26)
	for _, task := range tasks {
		cnts[task-'A'] += 1
	}
	maxCnt, leftCnt := 0, 0
	for _, cnt := range cnts {
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}
	fmt.Println("max cnt is ", maxCnt, cnts)
	for idx, cnt := range cnts {
		if cnt == maxCnt {
			leftCnt += 1
			fmt.Println(idx+1, cnt)
		}
	}
	return max(len(tasks), (maxCnt-1)*(n+1)+leftCnt)
}
