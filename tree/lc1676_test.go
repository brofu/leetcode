package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestorIV(t *testing.T) {
	type args struct {
		root  []int
		nodes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				root:  []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				nodes: []int{4, 7},
			},
			want: 2,
		},

		{
			name: "lc case 2",
			args: args{
				root:  []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				nodes: []int{1},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				root:  []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				nodes: []int{7, 6, 2, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			treeNodes := make([]*TreeNode, len(tt.args.nodes))
			for index, num := range tt.args.nodes {
				treeNodes[index] = GetTreeNodeFromTree(tree, num)
			}
			wantNode := GetTreeNodeFromTree(tree, tt.want)
			if got := lowestCommonAncestorIV(tree, treeNodes); !reflect.DeepEqual(got, wantNode) {
				t.Errorf("lowestCommonAncestorIV() = %v, want %v", got, tt.want)
			}
		})
	}
}
