package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
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
				p: &TreeNode{
					Val: 5,
				},
				q: &TreeNode{
					Val: 1,
				},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p: &TreeNode{
					Val: 5,
				},
				q: &TreeNode{
					Val: 4,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := lowestCommonAncestor(tree, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("case: %s, lowestCommonAncestor() = %v, want %v", tt.name, got.Val, tt.want)
			}
		})
	}
}
func Test_lowestCommonAncestorV2(t *testing.T) {
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
				p: &TreeNode{
					Val: 5,
				},
				q: &TreeNode{
					Val: 1,
				},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p: &TreeNode{
					Val: 5,
				},
				q: &TreeNode{
					Val: 4,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			got := lowestCommonAncestorV2(tree, tt.args.p, tt.args.q)
			if !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("case: %s, lowestCommonAncestorV2() = %v, want %v", tt.name, got.Val, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestorV2PV1(t *testing.T) {
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
				input: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p: &TreeNode{
					Val: 5,
				},
				q: &TreeNode{
					Val: 1,
				},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p: &TreeNode{
					Val: 5,
				},
				q: &TreeNode{
					Val: 4,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := lowestCommonAncestorV2PV1(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestorV2PV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestorPV1(t *testing.T) {
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
				input: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p: &TreeNode{
					Val: 5,
				},
				q: &TreeNode{
					Val: 1,
				},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p: &TreeNode{
					Val: 5,
				},
				q: &TreeNode{
					Val: 4,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := lowestCommonAncestorPV1(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestorPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
