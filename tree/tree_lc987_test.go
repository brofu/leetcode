package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_verticalTraversal(t *testing.T) {
	type args struct {
		input []int
		root  *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: [][]int{{9}, {3, 15}, {20}, {7}},
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
		{
			name: "lc case 3",
			args: args{
				input: []int{1, 2, 3, 4, 6, 5, 7},
			},
			want: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := verticalTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
