package tree

import "testing"

func Test_countNodes(t *testing.T) {
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
				root: []int{1, 2, 3, 4, 5, 6},
			},
			want: 6,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{},
			},
			want: 0,
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := countNodes(tree); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
