package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_balanceBST(t *testing.T) {
	type args struct {
		root  *TreeNode
		input []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{1, math.MaxInt, 2, math.MaxInt, math.MaxInt, math.MaxInt, 3, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 4},
			},
			want: [][]int{{2, 1, 3, 4}, {3, 1, 4, 2}},
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{2, 1, 3},
			},
			want: [][]int{{2, 1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.input, 0)
			got := balanceBST(tree)
			slice := GenerateSliceFromTreeNode(got)
			success := false
			for _, possible := range tt.want {
				if reflect.DeepEqual(slice, possible) {
					success = true
					continue
				}
			}
			if !success {
				t.Errorf("balanceBST() = %v, want %v", slice, tt.want)
			}
		})
	}
}
