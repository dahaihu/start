package leetcode

type iland struct {
	Points map[[2]int]struct{}
}

func newIland(key [2]int) *iland {
	return &iland{Points: map[[2]int]struct{}{key: struct{}{}}}
}

func maxAreaOfIsland(grid [][]int) int {
	mark := make(map[[2]int]*iland)
	var maxArea int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			var top, left *iland
			if i > 0 {
				top = mark[[2]int{i - 1, j}]
			}
			if j > 0 {
				left = mark[[2]int{i, j - 1}]
			}
			key := [2]int{i, j}
			var updatedLen int
			if top == nil && left == nil {
				mark[key] = newIland(key)
				updatedLen = 1
			} else if top == nil && left != nil {
				left.Points[key] = struct{}{}
				mark[key] = left
				updatedLen = len(left.Points)
			} else if top != nil && left == nil {
				top.Points[key] = struct{}{}
				mark[key] = top
				updatedLen = len(top.Points)
			} else {
				for key := range top.Points {
					left.Points[key] = struct{}{}
					mark[key] = left
				}
				left.Points[key] = struct{}{}
				mark[key] = left
				updatedLen = len(left.Points)
			}
			if updatedLen > maxArea {
				maxArea = updatedLen
			}
		}
	}
	return maxArea
}
