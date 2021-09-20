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
	if node.visited {
		return true
	}
	node.visited = true
	for _, child := range node.children {
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
		for _, item := range prerequisite {
			if _, ok := nodes[item]; !ok {
				nodes[item] = &Node{
					children:  nil,
					val:       item,
					notCircle: false,
					visited:   false,
				}
			}
		}
		child, parent := prerequisite[0], prerequisite[1]
		nodes[parent].children = append(nodes[parent].children, nodes[child])
	}
	for _, node := range nodes {
		if node.existCircle() {
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
	return len(courseChildren) == 0
}
