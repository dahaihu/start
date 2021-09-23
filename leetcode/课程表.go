package leetcode

type Node struct {
	val       int
	notCircle bool
	visited   bool
	children  []*Node
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
		for _, node := range prerequisite {
			if _, ok := nodes[node]; !ok {
				nodes[node] = &Node{val: node}
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
	courseChildren := make(map[int][]int)
	in := make(map[int]int)
	for _, prerequisite := range prerequisites {
		child, parent := prerequisite[0], prerequisite[1]
		in[child]++
		courseChildren[parent] = append(courseChildren[parent], child)
	}
	var queue []int
	for course := 0; course < numCourses; course++ {
		if in[course] == 0 {
			queue = append(queue, course)
			delete(in, course)
		}
	}
	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]
		for _, child := range courseChildren[course] {
			in[child]--
			if in[child] == 0 {
				queue = append(queue, child)
				delete(in, child)
			}
		}
	}
	return len(in) == 0
}
