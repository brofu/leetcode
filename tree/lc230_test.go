package tree

import (
	"math"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{3, 1, 4, math.MaxInt, 2},
				k:    1,
			},
			want: 1,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{5, 3, 6, 2, 4, math.MaxInt, math.MaxInt, 1},
				k:    3,
			},
			want: 3,
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{3, 1, 4, math.MaxInt, 2},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := kthSmallest(bst, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kthSmallestWrapper(t *testing.T) {
	type args struct {
		root  *TreeNode
		index int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := kthSmallestWrapper(tt.args.root, tt.args.index)
			if got != tt.want {
				t.Errorf("kthSmallestWrapper() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("kthSmallestWrapper() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
