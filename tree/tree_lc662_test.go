package tree

import (
	"math"
	"testing"
)

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		root  *TreeNode
		input []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		enabled bool
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{1, 3, 2, 5, 3, math.MaxInt, 9},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{1, 3, 2, 5, math.MaxInt, math.MaxInt, 9, 6, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 7},
			},
			want: 7,
		},
		{
			name: "lc case 3",
			args: args{
				input: []int{1, 3, 2, 5},
			},
			want: 2,
		},
		{
			name:    "lc case 4 - special case",
			enabled: false,
			args: args{
				input: []int{1, 1, 1, 1, 1, 1, 1, math.MaxInt, math.MaxInt, math.MaxInt, 1, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 2, 2, 2, 2, 2, 2, 2, math.MaxInt, 2, math.MaxInt, math.MaxInt, 2, math.MaxInt, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.enabled {
				return
			}
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := widthOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
