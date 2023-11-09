package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
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
				root: []int{1, math.MaxInt, 2, math.MaxInt, math.MaxInt, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := preorderTraversal(tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversalV2(t *testing.T) {
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
				root: []int{1, math.MaxInt, 2, math.MaxInt, math.MaxInt, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res = []int{} // need to clear the global var
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := preorderTraversalV2(tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
