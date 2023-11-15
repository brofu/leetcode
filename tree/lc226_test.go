package tree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{4, 2, 7, 1, 3, 6, 9},
			},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{2, 1, 3},
			},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			wantTree := GenerateBinaryTreeFromSlice(tt.want, 0)
			if got := invertTree(tree); !reflect.DeepEqual(got, wantTree) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
