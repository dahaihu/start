package leetcode

import "container/heap"

type priorityQueue struct {
	nums [][]int
}

func (p *priorityQueue) Less(i, j int) bool {
	return p.nums[i][2] < p.nums[j][2]
}

func (p *priorityQueue) Len() int {
	return len(p.nums)
}

func (p *priorityQueue) Swap(i, j int) {
	p.nums[i], p.nums[j] = p.nums[j], p.nums[i]
}

func (p *priorityQueue) Push(ele interface{}) {
	p.nums = append(p.nums, ele.([]int))
}

func (p *priorityQueue) Pop() interface{} {
	last := p.nums[len(p.nums)-1]
	p.nums = p.nums[:len(p.nums)-1]
	return last
}

func trapRainWater(heightMap [][]int) int {
	xLen, yLen := len(heightMap), len(heightMap[0])
	if xLen < 3 || yLen < 3 {
		return 0
	}
	visited := make([][]bool, xLen)
	for i := 0; i < xLen; i++ {
		visited[i] = make([]bool, yLen)
	}
	var pq priorityQueue
	for i := 0; i < xLen; i++ {
		for j := 0; j < yLen; j++ {
			if i == 0 || j == 0 || i == xLen-1 || j == yLen-1 {
				heap.Push(&pq, []int{i, j, heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}
	nexts := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var waters int
	for pq.Len() != 0 {
		cur := heap.Pop(&pq).([]int)
		for _, dir := range nexts {
			next := []int{cur[0] + dir[0], cur[1] + dir[1]}
			if next[0] >= 0 && next[0] < xLen &&
				next[1] >= 0 && next[1] < yLen &&
				!visited[next[0]][next[1]] {
				height := heightMap[next[0]][next[1]]
				temp := cur[2] - heightMap[next[0]][next[1]]
				if temp > 0 {
					waters += temp
					height = cur[2]
				}
				heap.Push(&pq, []int{next[0], next[1], height})
				visited[next[0]][next[1]] = true
			}
		}
	}
	return waters
}
