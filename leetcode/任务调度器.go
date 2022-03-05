package leetcode

import (
	"fmt"
	"sort"
)

/*
621. 任务调度器
给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。
然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
你需要计算完成所有任务所需要的 最短时间 。
示例 1：
输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B
     在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。

示例 2：
输入：tasks = ["A","A","A","B","B","B"], n = 0
输出：6
解释：在这种情况下，任何大小为 6 的排列都可以满足要求，因为 n = 0
["A","A","A","B","B","B"]
["A","B","A","B","A","B"]
["B","B","B","A","A","A"]
...
诸如此类

示例 3：
输入：tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
输出：16
解释：一种可能的解决方案是：
     A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> (待命) -> (待命) -> A -> (待命) -> (待命) -> A


提示：
1 <= task.length <= 104
tasks[i] 是大写英文字母
n 的取值范围为 [0, 100]
*/

func trim(nums []int) []int {
	endidx := len(nums) - 1
	for endidx >= 0 && nums[endidx] == 0 {
		endidx -= 1
	}
	return nums[:endidx+1]
}

func leastInterval(tasks []byte, n int) int {
	if n <= 1 {
		return len(tasks)
	}
	cnts := make([]int, 26)
	for _, task := range tasks {
		cnts[task-'A'] += 1
	}
	sort.Slice(cnts, func(i, j int) bool { return cnts[i] >= cnts[j] })
	cnts = trim(cnts)
	count := 0
	for {
		if len(cnts) == 0 {
			return count
		}
		if len(cnts) <= 1+n {
			count += (cnts[0]-1)*(n+1) + 1
			for i := 1; i < len(cnts); i++ {
				if cnts[i] == cnts[0] {
					count += 1
				} else {
					break
				}
			}
			return count
		}
		count += cnts[n] * (n + 1)
		for i := 0; i < n+1; i++ {
			cnts[i] -= cnts[n]
		}
		sort.Slice(cnts, func(i, j int) bool { return cnts[i] >= cnts[j] })
		cnts = trim(cnts)
	}
}

func insertAfter(nums [][]byte, cur, target int) {
	tmp := nums[cur]
	copy(nums[target+1:cur+1], nums[target:cur])
	nums[target] = tmp
}

func leastInterval2(tasks []byte, n int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i] < tasks[j] })
	mark := [][]byte{{tasks[0]}}
	for i := 1; i < len(tasks); i++ {
		if tasks[i] == tasks[i-1] {
			mark[len(mark)-1] = append(mark[len(mark)-1], tasks[i])
		} else {
			mark = append(mark, []byte{tasks[i]})
		}
	}
	sort.Slice(mark, func(i, j int) bool { return len(mark[i]) > len(mark[j]) })
	maxHeight := len(mark[0])
	maxCnt := 0
	var notMaxHeight int
	for i := 0; i < len(mark); i++ {
		if len(mark[i]) == maxHeight {
			maxCnt += 1
		} else {
			notMaxHeight = i
			break
		}
	}
	fmt.Println(mark)
	defer func() {
		fmt.Println(mark)
	}()
	for i := n + 1; i < len(mark); i++ {
		fmt.Println(i, mark)
		if len(mark[n]) >= maxHeight-1 {
			return len(tasks)
		}
		preHeight := maxHeight
		for len(mark[i]) != 0 {
			insertedLen := min(len(mark[i]), maxHeight-1-len(mark[n]))
			midMargin := min(preHeight-insertedLen, len(mark[n]))
			mid := make([]byte, len(mark[n][midMargin:]))
			copy(mid, mark[n][midMargin:])
			mark[n] = append(mark[n][:midMargin], mark[i][:insertedLen]...)
			mark[n] = append(mark[n], mid...)
			mark[i] = mark[i][insertedLen:]
			preHeight = preHeight - insertedLen
			// target := sort.Search(n, func(j int) bool { return len(mark[j]) < len(mark[n]) })
			// insertAfter(mark, n, target)
			if len(mark[n]) == maxHeight-1 {
				mark[notMaxHeight], mark[n] = mark[n], mark[notMaxHeight]
				notMaxHeight += 1
				if notMaxHeight > n {
					return len(tasks)
				}
			}
		}
		fmt.Println(i, mark)
	}
	fmt.Println(mark)
	return (maxHeight-1)*(n+1) + maxCnt
}

func leastIntervalBetter(tasks []byte, n int) int {
	cnts := make([]int, 26)
	for _, task := range tasks {
		cnts[task-'A']++
	}
	var maxCnt, equalMaxCnt int
	for _, cnt := range cnts {
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}
	for _, cnt := range cnts {
		if cnt == maxCnt {
			equalMaxCnt++
		}
	}

	return max(len(tasks), (maxCnt-1)*(n+1)+equalMaxCnt)
}
