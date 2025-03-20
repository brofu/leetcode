package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1. Iterative heapify",
			args: args{
				nums: []int{5, 3, 8, 4, 2, 7, 1, 10},
			},
			want: []int{10, 8, 7, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
			t.Log(tt.args.nums, tt.want)
		})
	}
}
