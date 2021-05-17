package leetcode

type Node struct {
	children  []*Node
	val       int
	notCircle bool
	visited   bool
}

func (node *Node) existCircle() bool {
	if node.notCircle {
		return false
	}
	node.visited = true
	for _, child := range node.children {
		if child.visited {
			return true
		}
		if child.existCircle() {
			return true
		}
	}
	node.visited = false
	node.notCircle = true
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	nodes := make(map[int]*Node, numCourses)

	for _, prerequisite := range prerequisites {
		for _, course := range prerequisite {
			if _, ok := nodes[course]; !ok {
				nodes[course] = &Node{val: course}
			}
		}
		child, parent := nodes[prerequisite[0]], nodes[prerequisite[1]]
		parent.children = append(parent.children, child)
	}
	for _, node := range nodes {
		if !node.notCircle && node.existCircle() {
			return false
		}
	}
	return true
}
