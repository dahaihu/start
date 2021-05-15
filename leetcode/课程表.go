package leetcode
type Node struct {
	parents      []*Node
	val          int
	visited      bool
	childChecked bool
}

func (node *Node) existCircle() bool {
	if node.childChecked {
		return false
	}
	node.visited = true
	for _, parent := range node.parents {
		if parent.visited {
			return true
		}
		if parent.existCircle() {
			return true
		}
	}
	node.visited = false
	node.childChecked = true
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	allCourses := make(map[int]*Node, numCourses)
	// 第一步，初始化树状图
	for _, prerequisite := range prerequisites {
		if prerequisite[0] == prerequisite[1] {
			return false
		}
		// init node when needed
		for _, course := range prerequisite {
			if _, ok := allCourses[course]; !ok {
				allCourses[course] = &Node{val: course}
			}
		}
		node := allCourses[prerequisite[0]]
		parent := allCourses[prerequisite[1]]
		node.parents = append(node.parents, parent)
	}
	// 第二步，找环
	for _, node := range allCourses {
		if !node.childChecked && node.existCircle() {
			return false
		}
	}
	return true
}

