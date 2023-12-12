package tree

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{4, 2, 7, 1, 3},
				val:  2,
			},
			want: []int{2, 1, 3},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{4, 2, 7, 1, 3},
				val:  5,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			wantTree := GenerateBinaryTreeFromSlice(tt.want, 0)
			if got := searchBST(tree, tt.args.val); !reflect.DeepEqual(got, wantTree) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
