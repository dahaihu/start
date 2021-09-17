package leetcode

import (
	"fmt"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{matrix: [][]int{
				{1, 2, 3},
				{1, 2, 0},
				{1, 9, 3},
			}},
		},
		{
			name: "test",
			args: args{matrix: [][]int{
				{1, 2, 3},
				{1, 0, 1},
				{1, 9, 1},
			}},
		},
		{
			name: "test",
			args: args{matrix: [][]int{
				{0, 2, 0},
				{1, 1, 1},
				{1, 9, 1},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			fmt.Println(tt.args.matrix)
		})
	}
}
