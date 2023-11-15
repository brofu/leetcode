package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_connect117(t *testing.T) {
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
				root: []int{1, 2, 3, 4, 5, math.MaxInt, 7},
			},
			want: []int{1, math.MaxInt, 2, 3, math.MaxInt, 4, 5, 7, math.MaxInt},
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
			got := connect117(tree)
			gotSlice := GenerateSlinceFromNode(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("connect117() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connectMulPointer(t *testing.T) {
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
				root: []int{1, 2, 3, 4, 5, math.MaxInt, 7},
			},
			want: []int{1, math.MaxInt, 2, 3, math.MaxInt, 4, 5, 7, math.MaxInt},
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
			got := connectMulPointer(tree)
			gotSlice := GenerateSlinceFromNode(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("connectMulPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
