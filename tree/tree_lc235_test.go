package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestorBST(t *testing.T) {
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
				root: []int{6, 2, 8, 0, 4, 7, 9, math.MaxInt, math.MaxInt, 3, 5},
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 8},
			},
			want: 6,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{6, 2, 8, 0, 4, 7, 9, math.MaxInt, math.MaxInt, 3, 5},
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 4},
			},
			want: 2,
		},

		{
			name: "lc case 3",
			args: args{
				root: []int{2, 1},
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 1},
			},
			want: 2,
		},
		{
			name: "lc case 4",
			args: args{
				root: []int{2, 1, 3},
				p:    &TreeNode{Val: 3},
				q:    &TreeNode{Val: 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := lowestCommonAncestorBST(tree, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestorBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestorBSTPV1(t *testing.T) {
	type args struct {
		root  *TreeNode
		input []int
		p     *TreeNode
		q     *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{6, 2, 8, 0, 4, 7, 9, math.MaxInt, math.MaxInt, 3, 5},
				p:     &TreeNode{Val: 2},
				q:     &TreeNode{Val: 8},
			},
			want: 6,
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{6, 2, 8, 0, 4, 7, 9, math.MaxInt, math.MaxInt, 3, 5},
				p:     &TreeNode{Val: 2},
				q:     &TreeNode{Val: 4},
			},
			want: 2,
		},

		{
			name: "lc case 3",
			args: args{
				input: []int{2, 1},
				p:     &TreeNode{Val: 2},
				q:     &TreeNode{Val: 1},
			},
			want: 2,
		},
		{
			name: "lc case 4",
			args: args{
				input: []int{2, 1, 3},
				p:     &TreeNode{Val: 3},
				q:     &TreeNode{Val: 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := lowestCommonAncestorBSTPV1(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestorBSTPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
