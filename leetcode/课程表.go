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


func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	courseChildren := make(map[int][]int, numCourses)
	// 这个地方为什么用切片，而不是字典，需要注意下
	in := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		child, parent := prerequisite[0], prerequisite[1]
		in[child] += 1
		courseChildren[parent] = append(courseChildren[parent], child)
	}
	var queue []int
	// 切片的遍历，包括可能的并没有出现在 prerequisites 中的课程，所以in这个变量最好是切片
	for course, inCount := range in {
		if inCount == 0 {
			queue = append(queue, course)
		}
	}

	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]
		for _, child := range courseChildren[course] {
			in[child] -= 1
			if in[child] == 0 {
				queue = append(queue, child)
			}
		}
		delete(courseChildren, course)
	}
	if len(courseChildren) == 0 {
		return true
	}
	return false
}