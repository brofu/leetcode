package tree

import (
	"reflect"
	"testing"
)

func Test_constructFromPrePostSubTaskPV2(t *testing.T) {
	type args struct {
		preorder  []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case 1",
			args: args{
				preorder:  []int{2, 1, 3},
				postorder: []int{3, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructFromPrePostSubTaskPV2(tt.args.preorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructFromPrePostSubTaskPV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_constructFromPrePostSubTaskPV3(t *testing.T) {
	type args struct {
		preorder  []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case 1",
			args: args{
				preorder:  []int{1, 2, 4, 5, 3, 6, 7},
				postorder: []int{4, 5, 2, 6, 7, 3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructFromPrePostSubTaskPV3(tt.args.preorder, tt.args.postorder)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructFromPrePostSubTaskPV3() = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}
