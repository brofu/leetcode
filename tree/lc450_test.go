package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root []int
		key  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{5, 3, 6, 2, 4, math.MaxInt, 7},
				key:  3,
			},
			want: []int{5, 4, 6, 2, math.MaxInt, math.MaxInt, 7},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{5, 3, 6, 2, 4, math.MaxInt, 7},
				key:  0,
			},
			want: []int{5, 3, 6, 2, 4, math.MaxInt, 7},
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{},
				key:  0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			wantTree := GenerateBinaryTreeFromSlice(tt.want, 0)
			if got := deleteNode(tree, tt.args.key); !reflect.DeepEqual(got, wantTree) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
