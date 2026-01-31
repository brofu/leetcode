package stack

import (
	"testing"

	"github.com/brofu/leetcode/list"
	"github.com/stretchr/testify/assert"
)

func Test_reorderList(t *testing.T) {
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
				head: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 5, 2, 4, 3},
		},
		{
			name: "case 2",
			args: args{
				head: []int{1, 2, 3, 4},
			},
			want: []int{1, 4, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.ConstructListNodeFromSlice(tt.args.head)
			reorderList(l)
			got := list.GetSliceFromList(l)
			t.Log("got", got)
			t.Log("want", tt.want)
			assert.Equal(t, tt.want, got)
		})
	}
}
