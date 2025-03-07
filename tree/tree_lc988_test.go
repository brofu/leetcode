package tree

import (
	"math"
	"testing"
)

func Test_smallestFromLeaf(t *testing.T) {
	type args struct {
		root    *TreeNode
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal case 1",
			args: args{
				numbers: []int{2, 2, 1, math.MaxInt, 1, 0, math.MaxInt, 0},
			},
			want: "abc",
		},
		{
			name: "normal case 2",
			args: args{
				numbers: []int{0, math.MaxInt, 1},
			},
			want: "ba",
		},
		{
			name: "normal case 3",
			args: args{
				numbers: []int{4, 0, 1, 1},
			},
			want: "bae",
		},
		{
			name: "normal case 4",
			args: args{
				numbers: []int{25, 1, math.MaxInt, 0, 0, 1, math.MaxInt, math.MaxInt, math.MaxInt, 0},
			},
			want: "ababz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSliceBFS(tt.args.numbers)
			if got := smallestFromLeaf(tt.args.root); got != tt.want {
				t.Errorf("smallestFromLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestFromLeafDFS(t *testing.T) {
	type args struct {
		root    *TreeNode
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal case 1",
			args: args{
				numbers: []int{2, 2, 1, math.MaxInt, 1, 0, math.MaxInt, 0},
			},
			want: "abc",
		},
		{
			name: "normal case 2",
			args: args{
				numbers: []int{0, math.MaxInt, 1},
			},
			want: "ba",
		},
		{
			name: "normal case 3",
			args: args{
				numbers: []int{4, 0, 1, 1},
			},
			want: "bae",
		},
		{
			name: "normal case 4",
			args: args{
				numbers: []int{25, 1, math.MaxInt, 0, 0, 1, math.MaxInt, math.MaxInt, math.MaxInt, 0},
			},
			want: "ababz",
		},
		{
			name: "normal case 5",
			args: args{
				numbers: []int{0, 1, math.MaxInt, math.MaxInt, 2, 6, 3, math.MaxInt, math.MaxInt, math.MaxInt, 4, math.MaxInt, 5},
			},
			want: "fedcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.args.root = GenerateBinaryTreeFromSliceBFS(tt.args.numbers)
			if got := smallestFromLeafDFS(tt.args.root); got != tt.want {
				t.Errorf("smallestFromLeafDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
