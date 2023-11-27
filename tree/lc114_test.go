package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{1, 2, 5, 3, 4, math.MaxInt, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{0},
			},
			want: []int{0},
		},
		{
			name: "lc case 4",
			args: args{
				root: []int{1, 2, math.MaxInt, 3, 4, math.MaxInt, math.MaxInt, 5},
			},
			want: []int{1, 2, 3, 5, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			flatten(tree)
			gotSlice := GenerateSliceFromTreeNode(tree)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("case name: %s, flatten() = %v, want %v", tt.name, gotSlice, tt.want)
			}
		})
	}
}

func Test_flattenTraverse(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenTraverse(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flattenV2(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{1, 2, 5, 3, 4, math.MaxInt, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		//		{
		//			name: "lc case 2",
		//			args: args{
		//				root: []int{},
		//			},
		//			want: []int{},
		//		},
		{
			name: "lc case 3",
			args: args{
				root: []int{0},
			},
			want: []int{0},
		},
		{
			name: "lc case 4",
			args: args{
				root: []int{1, 2, math.MaxInt, 3, 4, math.MaxInt, math.MaxInt, 5},
			},
			want: []int{1, 2, 3, 5, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			flattenV2(&tree)
			gotSlice := GenerateSliceFromTreeNode(tree)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("case name: %s, flattenV2() = %v, want %v", tt.name, gotSlice, tt.want)
			}
		})
	}
}
