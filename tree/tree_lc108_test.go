package tree

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"lc case 1",
			args{
				[]int{-10, -3, 0, 5, 9},
			},
			[][]int{{0, -3, 9, -10, 5}, {0, -10, 5, -3, 9}},
		},
		{
			"lc case 2",
			args{
				[]int{1, 3},
			},
			[][]int{{1, 3}, {3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.args.nums)
			slice := GenerateSliceFromTreeNode(got)
			success := false
			for _, possible := range tt.want {
				if reflect.DeepEqual(slice, possible) {
					success = true
					continue
				}
			}
			if !success {
				t.Errorf("sortedArrayToBST() = %v, want %v", slice, tt.want)
			}
		})
	}
}
