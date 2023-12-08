package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_convertBST(t *testing.T) {
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
				root: []int{4, 1, 6, 0, 2, 5, 7, math.MaxInt, math.MaxInt, math.MaxInt, 3, math.MaxInt, math.MaxInt, math.MaxInt, 8},
			},
			want: []int{30, 36, 21, 36, 35, 26, 15, 33, 8},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{0, math.MaxInt, 1},
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			got := convertBST(tree)
			gotSlice := GenerateSliceFromTreeNode(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("convertBST() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}

func Test_convertBSTWrapper(t *testing.T) {
	type args struct {
		root  *TreeNode
		count int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := convertBSTWrapper(tt.args.root, tt.args.count); gotRes != tt.wantRes {
				t.Errorf("convertBSTWrapper() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_convertBSTV2(t *testing.T) {
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
				root: []int{4, 1, 6, 0, 2, 5, 7, math.MaxInt, math.MaxInt, math.MaxInt, 3, math.MaxInt, math.MaxInt, math.MaxInt, 8},
			},
			want: []int{30, 36, 21, 36, 35, 26, 15, 33, 8},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{0, math.MaxInt, 1},
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			got := convertBSTV2(tree)
			gotSlice := GenerateSliceFromTreeNode(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("convertBSTV2() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
