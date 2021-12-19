package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name string
	Age  int
}

func binarySearchLeftPosition(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func editDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	mark := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		mark[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		mark[i][0] = i
	}
	for i := 1; i <= n; i++ {
		mark[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				mark[i][j] = mark[i-1][j-1]
			} else {
				mark[i][j] = min(
					mark[i-1][j-1],
					min(mark[i-1][j], mark[i][j-1]),
				) + 1
			}
		}
	}
	return mark[m][n]
}

func splitArray(nums1, nums2 []int) (int, int) {
	// nums1Len < nums2Len
	nums1Len, nums2Len := len(nums1), len(nums2)
	count := (nums1Len + nums2Len + 1) / 2
	// 思考，进步，远远的超过
	left, right := 0, nums1Len
	for left < right {
		nums1Idx := (right-left)/2 + left
		nums2Idx := count - nums1Idx
		if nums1[nums1Idx] >= nums2[nums2Idx-1] {
			right = nums1Idx
		} else {
			left = nums1Idx + 1
		}
	}
	return left, count - left
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func arrayIdx(array []int, idx int, defaultVal int) int {
	if idx >= 0 && idx < len(array) {
		return array[idx]
	}
	return defaultVal
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	nums1Len, nums2Len := len(nums1), len(nums2)
	if nums1Len > nums2Len {
		return findMedianSortedArrays(nums2, nums1)
	}
	nums1Idx, nums2Idx := splitArray(nums1, nums2)
	leftVal := max(arrayIdx(nums1, nums1Idx-1, math.MinInt64),
		arrayIdx(nums2, nums2Idx-1, math.MinInt64))
	if (nums1Len+nums2Len)%2 == 1 {
		return float64(leftVal)
	}
	rightVal := min(arrayIdx(nums1, nums1Idx, math.MaxInt64),
		arrayIdx(nums2, nums2Idx, math.MaxInt64))
	return (float64(leftVal) + float64(rightVal)) / 2
}

func trap(height []int) int {
	waters := 0
	indexes := []int{}
	for idx, h := range height {
		preHeight := 0
		for len(indexes) != 0 && h >= height[indexes[len(indexes)-1]] {
			preIdx := indexes[len(indexes)-1]
			waters += (idx - preIdx - 1) * (height[preIdx] - preHeight)
			preHeight = height[preIdx]
			indexes = indexes[:len(indexes)-1]
		}
		if len(indexes) != 0 {
			preIdx := indexes[len(indexes)-1]
			waters += (idx - preIdx - 1) * (h - preHeight)
		}
		indexes = append(indexes, idx)
	}
	return waters
}

type element struct {
	Index  int
	Height int
}

func largestRectangleArea(heights []int) int {
	var mx int
	mark := make([]*element, 0, len(heights))
	for idx, height := range heights {
		if len(mark) == 0 || mark[len(mark)-1].Height < height {
			mark = append(mark, &element{
				Index:  idx,
				Height: height,
			})
			continue
		}
		var preIdx int
		for len(mark) != 0 && mark[len(mark)-1].Height >= height {
			lastItem := mark[len(mark)-1]
			rectangleArea := (idx - lastItem.Index) * lastItem.Height
			if rectangleArea > mx {
				mx = rectangleArea
			}
			preIdx = lastItem.Index
			mark = mark[:len(mark)-1]
		}
		mark = append(mark, &element{
			Index:  preIdx,
			Height: height,
		})
	}
	for _, ele := range mark {
		if tmp := ele.Height * (len(heights) - ele.Index); tmp > mx {
			mx = tmp
		}
	}
	return mx
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) GetVal() int {
	if n == nil {
		return 0
	}
	return n.Val
}

func (n *ListNode) GetNext() *ListNode {
	if n == nil {
		return nil
	}
	return n.Next
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	return _addTwoNumbers(l1, l2, 0)

}

func _addTwoNumbers(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	var node *ListNode
	if l1 != nil {
		node = l1
	} else if l2 != nil {
		node = l2
	} else {
		node = new(ListNode)
	}

	node.Val = l1.GetVal() + l2.GetVal() + carry

	if node.Val >= 10 {
		node.Val -= 10
		carry = 1
	} else {
		carry = 0
	}

	node.Next = _addTwoNumbers(l1.GetNext(), l2.GetNext(), carry)
	return node
}

func main() {
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{Val: 5}
	node5 := &ListNode{Val: 6}
	node6 := &ListNode{Val: 4}
	node4.Next = node5
	node5.Next = node6

	head := addTwoNumbers(node1, node4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
