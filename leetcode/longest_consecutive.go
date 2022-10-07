package leetcode

import "fmt"

func longestConsecutive(nums []int) int {
	mark := make(map[int]int)
	res := 0
	for _, num := range nums {
		if _, ok := mark[num]; ok {
			continue
		}
		left := mark[num-1]
		right := mark[num+1]
		cur := left + right + 1
		if cur > res {
			res = cur
		}
		mark[num] = cur
		mark[num-left] = cur
		mark[num+right] = cur
	}
	return res
}

var placeHolder struct{}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(t T) {
	s[t] = placeHolder
}

func (s Set[T]) Remove(t T) {
	delete(s, t)
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) Merge(other Set[T]) {
	for key := range other {
		s.Add(key)
	}
}

func (s Set[T]) Len() int {
	return len(s)
}

func root(ancestor map[int]int, target int) int {
	for {
		parent := ancestor[target]
		if parent == target {
			return parent
		}
		target = parent
	}
}

func longestConsecutiveUsingSet(nums []int) int {
	children := make(map[int]Set[int])
	for _, num := range nums {
		s := NewSet[int]()
		s.Add(num)
		children[num] = s
	}
	ancestor := make(map[int]int)
	var longest int
	for _, num := range nums {
		if _, visited := ancestor[num]; visited {
			continue
		}
		var curRoot int
		{
			left := num - 1
			_, leftOK := ancestor[left]
			if leftOK {
				curRoot = root(ancestor, left)
				children[curRoot].Add(num)
				delete(children, num)
			} else {
				curRoot = num
			}
		}
		ancestor[num] = curRoot
		{
			right := num + 1
			_, rightOK := ancestor[right]
			if rightOK {
				children[curRoot].Merge(children[right])
				delete(children, right)
				ancestor[right] = curRoot
			}
		}
		if length := children[curRoot].Len(); length > longest {
			longest = length
		}
	}
	fmt.Println(ancestor)
	fmt.Println(children)
	return longest
}
