package tree

import (
	"math"
	"reflect"
	"testing"
)

func TestBreadthFirstTraverse(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, math.MaxInt, 3, math.MaxInt, math.MaxInt, 4, 5},
			},
			want: []int{1, 3, 4, 5},
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, 3, 4, 5, math.MaxInt, math.MaxInt},
			},
			want: []int{1, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := BreadthFirstTraverse(tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreadthFirstTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBreadthFirstTraverseRecursive(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, math.MaxInt, 3, math.MaxInt, math.MaxInt, 4, 5},
			},
			want: []int{1, 3, 4, 5},
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, 3, 4, 5, math.MaxInt, math.MaxInt},
			},
			want: []int{1, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateBinaryTreeFromSlice(tt.args.root, 0)
			if got := BreadthFirstTraverseRecursive(tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreadthFirstTraverseRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
