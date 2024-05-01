package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
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
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{1},
			},
			want: [][]int{{1}},
		},
		{
			name: "lc case 3",
			args: args{
				input: []int{},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zigzagLevelOrderDFS(t *testing.T) {
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
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{1},
			},
			want: [][]int{{1}},
		},
		{
			name: "lc case 3",
			args: args{
				input: []int{},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := zigzagLevelOrderDFS(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrderDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
