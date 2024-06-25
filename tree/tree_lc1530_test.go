package tree

import (
	"math"
	"testing"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		root     *TreeNode
		input    []int
		distance int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				input:    []int{1, 2, 3, math.MaxInt, 4},
				distance: 3,
			},
			want: 1,
		},
		{
			name: "lc case 2",
			args: args{
				input:    []int{1, 2, 3, 4, 5, 6, 7},
				distance: 3,
			},
			want: 2,
		},
		{
			name: "lc case 3",
			args: args{
				input:    []int{7, 1, 4, 6, math.MaxInt, 5, 3, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 2},
				distance: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := countPairs(tt.args.root, tt.args.distance); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
