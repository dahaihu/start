package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	// start 表示起点， end 表示终点
	// start 不包括在内，end 包括在内
	end, start := 0, len(gas)-1
	res := gas[0] - cost[0]
	for end < start {
		if res < 0 {
			res += gas[start] - cost[start]
			start -= 1
		} else {
			end += 1
			res += gas[end] - cost[end]
		}
	}
	if res < 0 {
		return -1
	}
	return (start + 1) % len(gas)
}
