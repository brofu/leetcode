package tree

import (
	"math"
	"testing"
)

func Test_pseudoPalindromicPaths(t *testing.T) {
	type args struct {
		root *TreeNode
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{2, 3, 1, 3, 1, math.MaxInt, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateTreeNodeFromSliceBFS(tt.args.nums)

			got := pseudoPalindromicPaths(tt.args.root)
			if got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
