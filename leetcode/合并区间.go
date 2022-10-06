package leetcode

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"sort"
)

type Interval []int

func (i Interval) First() int {
	return i[0]
}

func (i Interval) Last() int {
	return i[1]
}

func (i Interval) SetFirst(first int) {
	i[0] = first
}

func (i Interval) SetLast(last int) {
	i[1] = last
}

type Intervals []Interval

func (is Intervals) Len() int {
	return len(is)
}

func (is Intervals) Less(i, j int) bool {
	switch {
	case is[i].First() == is[j].First():
		return is[i].Last() <= is[j].Last()
	default:
		return is[i].First() < is[j].First()
	}
}

func (is Intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func NewIntervals(intervals [][]int) Intervals {
	out := make(Intervals, 0, len(intervals))
	for _, interval := range intervals {
		out = append(out, NewInterval(interval))
	}
	return out
}

func NewInterval(interval []int) Interval {
	return Interval{interval[0], interval[1]}
}

func merge(intervals Intervals) Intervals {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(intervals)
	result := make(Intervals, 1)
	result[0] = intervals[0]
	for i := 1; i < len(intervals); i++ {
		pre, cur := result[len(result)-1], intervals[i]
		switch {
		case pre.Last() < cur.First():
			result = append(result, cur)
		case pre.Last() < cur.Last():
			pre.SetLast(cur.Last())
		}
	}
	return result
}
