package leetcode

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSearchRange(t *testing.T) {
	assert.Equal(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	assert.Equal(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	assert.Equal(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 1), []int{-1, -1})
	assert.Equal(t, searchRange([]int{}, 0), []int{-1, -1})
}
