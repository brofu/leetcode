package tree

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: []int{3, 9, 20, 15, 7},
		},
		{
			name: "lc case 2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := buildTree(tt.args.preorder, tt.args.inorder)
			gotSlice := GenerateSliceFromTreeNode(tree)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("buildTree() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
