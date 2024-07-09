package tree

import (
	"reflect"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 3,
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: [][]int{{1}},
		},

		{
			name: "lc case 3",
			args: args{
				n: 4,
			},
			want: [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 4, 2, 3}, {1, 4, 3, 2}, {2, 1, 3, 4}, {2, 1, 4, 3}, {3, 1, 4, 2}, {3, 2, 4, 1}, {4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 3, 1, 2}, {4, 3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTrees(tt.args.n)
			gotSlices := make([][]int, len(got))

			for i, tree := range got {
				gotSlices[i] = GenerateSliceFromTreeNode(tree)
			}
			if !reflect.DeepEqual(gotSlices, tt.want) {
				t.Errorf("generateTrees() = %v, want %v", gotSlices, tt.want)
			}
		})
	}
}

func Test_generateTreesPV1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 3,
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: [][]int{{1}},
		},

		{
			name: "lc case 3",
			args: args{
				n: 4,
			},
			want: [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 4, 2, 3}, {1, 4, 3, 2}, {2, 1, 3, 4}, {2, 1, 4, 3}, {3, 1, 4, 2}, {3, 2, 4, 1}, {4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 3, 1, 2}, {4, 3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTreesPV1(tt.args.n)
			gotSlices := make([][]int, len(got))

			for i, tree := range got {
				gotSlices[i] = GenerateSliceFromTreeNode(tree)
			}
			if !reflect.DeepEqual(gotSlices, tt.want) {
				t.Errorf("generateTreesPV1() = %v, want %v", gotSlices, tt.want)
			}
		})
	}
}
