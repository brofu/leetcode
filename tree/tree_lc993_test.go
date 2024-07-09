package tree

import (
	"math"
	"testing"
)

func Test_isCousins(t *testing.T) {
	type args struct {
		root  *TreeNode
		input []int
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{1, 2, 3, 4},
				x:     4,
				y:     3,
			},
			want: false,
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{1, 2, 3, math.MaxInt, 4, math.MaxInt, 5},
				x:     5,
				y:     4,
			},
			want: true,
		},
		{
			name: "lc case 3",
			args: args{
				input: []int{1, 2, 3, math.MaxInt, 4},
				x:     2,
				y:     3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
