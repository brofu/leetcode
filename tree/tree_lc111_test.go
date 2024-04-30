package tree

import (
	"math"
	"testing"
)

func Test_minDepth(t *testing.T) {
	type args struct {
		root  *TreeNode
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: 2,
		},

		{
			name: "lc case 2",
			args: args{
				input: []int{2, math.MaxInt, 3, math.MaxInt, math.MaxInt, math.MaxInt, 4, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 5, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 6},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepthV2(t *testing.T) {
	type args struct {
		root  *TreeNode
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: 2,
		},

		{
			name: "lc case 2",
			args: args{
				input: []int{2, math.MaxInt, 3, math.MaxInt, math.MaxInt, math.MaxInt, 4, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 5, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 6},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := minDepthV2(tt.args.root); got != tt.want {
				t.Errorf("minDepthV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
