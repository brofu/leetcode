package tree_common

import (
	"testing"

	"github.com/brofu/leetcode/common"
	"github.com/stretchr/testify/assert"
)

//wrapInt implements the common.Comparable interface

type wrapInt struct {
	number int
}

func (wi *wrapInt) CompareTo(comparbale common.Comparable) int {
	wrapIntNumber, _ := comparbale.(*wrapInt)
	if wi.number > wrapIntNumber.number {
		return 1
	}
	if wi.number == wrapIntNumber.number {
		return 0
	}
	return -1
}

func TestNewMaxPQ(t *testing.T) {
	type args struct {
		elements []common.Comparable
	}
	tests := []struct {
		name string
		args args
		want []common.Comparable
	}{
		{
			name: "construct PQ. normal case",
			args: args{
				elements: []common.Comparable{
					&wrapInt{number: 5},
					&wrapInt{number: 1},
					&wrapInt{number: 3},
					&wrapInt{number: 4},
				},
			},
			want: []common.Comparable{
				nil,
				&wrapInt{number: 5},
				&wrapInt{number: 4},
				&wrapInt{number: 3},
				&wrapInt{number: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// new
			got := NewMaxPQ(tt.args.elements)
			t.Log("flag1")
			assert.Equal(t, len(tt.want), len(got.Queue()))
			for i := 1; i < len(tt.want); i += 1 {
				wantWrapInt, ok := tt.want[i].(*wrapInt)
				assert.Equal(t, true, ok)
				gotWrapInt, ok := got.Queue()[i].(*wrapInt)
				assert.Equal(t, true, ok)
				assert.Equal(t, wantWrapInt.number, gotWrapInt.number)
			}

			//
			for i := 1; i < len(tt.want); i += 1 {
				gotComparable := got.Pop()
				gotWrapInt, ok := gotComparable.(*wrapInt)
				assert.Equal(t, true, ok)
				wantWrapInt, ok := tt.want[i].(*wrapInt)
				assert.Equal(t, true, ok)
				assert.Equal(t, wantWrapInt.number, gotWrapInt.number)
			}
		})
	}
}
