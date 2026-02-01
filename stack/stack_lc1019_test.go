package stack

import (
	"reflect"
	"testing"

	"github.com/brofu/leetcode/list"
)

func Test_nextLargerNodes(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				head: []int{2, 1, 5},
			},
			want: []int{5, 5, 0},
		},
		{
			name: "case 2",
			args: args{
				head: []int{2, 7, 4, 3, 5},
			},
			want: []int{7, 0, 5, 5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.ConstructListNodeFromSlice(tt.args.head)
			got := nextLargerNodes(l)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
