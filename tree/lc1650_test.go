package tree

import (
	"math"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestorIII(t *testing.T) {
	type args struct {
		root []int
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p:    5,
				q:    1,
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p:    5,
				q:    4,
			},
			want: 5,
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{1, 2},
				p:    1,
				q:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateTreeNodeWithParentFromSlice(tt.args.root, 0, nil)
			p := GetNodeWithParentFromTree(tree, tt.args.p)
			q := GetNodeWithParentFromTree(tree, tt.args.q)
			wantNode := GetNodeWithParentFromTree(tree, tt.want)
			if got := lowestCommonAncestorIII(p, q); !reflect.DeepEqual(got, wantNode) {
				t.Errorf("lowestCommonAncestorIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestorIIIV2(t *testing.T) {
	type args struct {
		root []int
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p:    5,
				q:    1,
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, math.MaxInt, math.MaxInt, 7, 4},
				p:    5,
				q:    4,
			},
			want: 5,
		},
		{
			name: "lc case 3",
			args: args{
				root: []int{1, 2},
				p:    1,
				q:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := GenerateTreeNodeWithParentFromSlice(tt.args.root, 0, nil)
			p := GetNodeWithParentFromTree(tree, tt.args.p)
			q := GetNodeWithParentFromTree(tree, tt.args.q)
			wantNode := GetNodeWithParentFromTree(tree, tt.want)
			if got := lowestCommonAncestorIIIV2(p, q); !reflect.DeepEqual(got, wantNode) {
				t.Errorf("lowestCommonAncestorIIIV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
