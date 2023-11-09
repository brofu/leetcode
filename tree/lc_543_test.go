package tree

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
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
				root: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := diameterOfBinaryTree(tree); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diameterOfBinaryTreeV2(t *testing.T) {
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
				root: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := diameterOfBinaryTreeV2(tree); got != tt.want {
				t.Errorf("diameterOfBinaryTreeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
