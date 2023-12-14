package tree

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestorII(t *testing.T) {
	type args struct {
		root []int
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p:    &TreeNode{Val: 5},
				q:    &TreeNode{Val: 1},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p:    &TreeNode{Val: 5},
				q:    &TreeNode{Val: 4},
			},
			want: 5,
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p:    &TreeNode{Val: 5},
				q:    &TreeNode{Val: 10},
			},
			want: math.MaxInt,
		},

		{
			name: "lc case 4",
			args: args{
				root: []int{17, 13, math.MaxInt, 15, math.MaxInt, math.MaxInt, math.MaxInt, 18, 5, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 4, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 9},
				p:    &TreeNode{Val: 5},
				q:    &TreeNode{Val: 9},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			got := lowestCommonAncestorII(tree, tt.args.p, tt.args.q)
			if tt.want == math.MaxInt {
				if !assert.Equal(t, *new(*TreeNode), got) {
					t.Errorf("lowestCommonAncestorII() = %v, want %v", got, tt.want)
				}
			} else if !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestorII() = %v, want %v", got, tt.want)
			}
		})
	}
}
