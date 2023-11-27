package tree

import (
	"reflect"
	"testing"
)

func Test_constructFromPrePost(t *testing.T) {
	type args struct {
		preorder  []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				preorder:  []int{1, 2, 4, 5, 3, 6, 7},
				postorder: []int{4, 5, 2, 6, 7, 3, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "lc case 2",
			args: args{
				preorder:  []int{1},
				postorder: []int{1},
			},
			want: []int{1},
		},
		{
			name: "lc case 3",
			args: args{
				preorder:  []int{2, 1, 3},
				postorder: []int{3, 1, 2},
			},
			want: []int{2, 1, 3},
		},
		{
			name: "lc case 4",
			args: args{
				preorder:  []int{3, 4, 2, 1},
				postorder: []int{2, 1, 4, 3},
			},
			want: []int{3, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructFromPrePost(tt.args.preorder, tt.args.postorder)
			gotSlice := GenerateSliceFromTreeNode(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("constructFromPrePost() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
