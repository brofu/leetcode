package tree

import (
	"math"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{2, 1, 3},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{5, 1, 4, math.MaxInt, math.MaxInt, 6},
			},
			want: false,
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{2, 2, 2},
			},
			want: false,
		},
		{
			name: "lc case 4",
			args: args{
				root: []int{5, 4, 6, math.MaxInt, math.MaxInt, 3, 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := isValidBST(tree); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
