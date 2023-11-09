package tree

import (
	"math"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{1, math.MaxInt, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := maxDepth(tree); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
