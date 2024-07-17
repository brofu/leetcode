package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_verticalOrder(t *testing.T) {
	type args struct {
		root  *TreeNode
		input []int
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
				input: []int{3, 9, 8, 4, 0, 1, 7},
			},
			want: [][]int{{4}, {9}, {3, 0, 1}, {8}, {7}},
		},
		{
			name: "lc case 3",
			args: args{
				input: []int{3, 9, 8, 4, 0, 1, 7, math.MaxInt, math.MaxInt, 5, math.MaxInt, math.MaxInt, 2},
			},
			want: [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}},
		},
		{
			name: "lc case 4",
			args: args{
				input: []int{3, 9, 8, 4, 0, 1, 7, math.MaxInt, math.MaxInt, math.MaxInt, 2, 5},
			},
			want: [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}},
		},
		{
			name: "lc case 5",
			args: args{
				input: []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := verticalOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verticalOrderWrong(t *testing.T) {
	type args struct {
		root  *TreeNode
		input []int
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
				input: []int{3, 9, 8, 4, 0, 1, 7},
			},
			want: [][]int{{4}, {9}, {3, 0, 1}, {8}, {7}},
		},
		{
			name: "lc case 3",
			args: args{
				input: []int{3, 9, 8, 4, 0, 1, 7, math.MaxInt, math.MaxInt, math.MaxInt, 2, 5},
			},
			want: [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalOrderWrong(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalOrderWrong() = %v, want %v", got, tt.want)
			}
		})
	}
}
