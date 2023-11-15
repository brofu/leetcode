package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
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
				root: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: []int{1, math.MaxInt, 2, 3, math.MaxInt, 4, 5, 6, 7, math.MaxInt},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateNodeFromSlice(tt.args.root, 0)
			got := connect(tree)
			gotSlice := GenerateSlinceFromNode(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectTraverse(t *testing.T) {
	type args struct {
		nodeLeft  *Node
		nodeRight *Node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connectTraverse(tt.args.nodeLeft, tt.args.nodeRight)
		})
	}
}

func Test_connectBFS(t *testing.T) {
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
				root: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: []int{1, math.MaxInt, 2, 3, math.MaxInt, 4, 5, 6, 7, math.MaxInt},
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateNodeFromSlice(tt.args.root, 0)
			got := connectBFS(tree)
			gotSlice := GenerateSlinceFromNode(got)

			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("connectBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
