package tree

import "testing"

func TestQuery(t *testing.T) {
	type args struct {
		root   *TreeNode
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"self test 1",
			args{
				input:  []int{0, 1, 2, 3, 4},
				target: 5,
			},
			false,
		},
		{
			"self test 2",
			args{
				input:  []int{0, 1, 2, 3, 4, 5},
				target: 5,
			},
			true,
		},
		{
			"self test 3",
			args{
				input:  []int{0, 1, 2, 3, 4, 5},
				target: 2,
			},
			true,
		},
		{
			"self test 4",
			args{
				input:  []int{0, 1, 2, 3, 4, 5},
				target: 8,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.root = GenerateBinaryTreeFromSlice(tt.args.input, 0)
			if got := Query(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
