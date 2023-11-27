package tree

import (
	"reflect"
	"testing"
)

func Test_buildTreeV2(t *testing.T) {
	type args struct {
		postorder []int
		inorder   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				inorder:   []int{9, 3, 15, 20, 7},
				postorder: []int{9, 15, 7, 20, 3},
			},
			want: []int{3, 9, 20, 15, 7},
		},
		{
			name: "lc case 2",
			args: args{
				postorder: []int{-1},
				inorder:   []int{-1},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := buildTreeV2(tt.args.inorder, tt.args.postorder)
			gotSlice := GenerateSliceFromTreeNode(tree)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("buildTreeV2() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
