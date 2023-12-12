package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
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
				val:  5,
			},
			want: []int{4, 2, 7, 1, 3, 5},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{40, 20, 60, 10, 30, 50, 70},
				val:  25,
			},
			want: []int{40, 20, 60, 10, 30, 50, 70, math.MaxInt, math.MaxInt, 25},
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{4, 2, 7, 1, 3, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt},
				val:  5,
			},
			want: []int{4, 2, 7, 1, 3, 5},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			wantTree := GenerateBinaryTreeFromSlice(tt.want, 0)
			if got := insertIntoBST(tree, tt.args.val); !reflect.DeepEqual(got, wantTree) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
