package tree

import (
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{3, 2, 1, 6, 0, 5},
			},
			want: []int{6, 3, 5, 2, 0, 1},
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructMaximumBinaryTree(tt.args.nums)
			gotSlice := GenerateSliceFromTreeNode(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
