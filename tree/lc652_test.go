package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{1, 2, 3, 4, math.MaxInt, 2, 4, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 4},
			},
			want: [][]int{{4}, {2, 4}},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{2, 1, 1},
			},
			want: [][]int{{1}},
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{2, 2, 2, 3, math.MaxInt, 3},
			},
			want: [][]int{{3}, {2, 3}},
		},
		{
			name: "lc case 4",
			args: args{
				root: []int{0, 0, 0, 0, math.MaxInt, math.MaxInt, 0, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 0, 0},
			},
			want: [][]int{{0}},
		},
		{
			name: "lc case 5",
			args: args{
				root: []int{0, 0, 0, 0, math.MaxInt, math.MaxInt, 0, 0, 0, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 0, 0},
			},
			want: [][]int{{0}, {0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			got := findDuplicateSubtrees(tree)
			slices := make([][]int, len(got))
			for index, tree := range got {
				slices[index] = GenerateSliceFromTreeNode(tree)
			}
			if !reflect.DeepEqual(slices, tt.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want %v", slices, tt.want)
			}
		})
	}
}
