package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPQBasedOnBH(t *testing.T) {
	type args struct {
		size int
		f    func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want *PQBasedOnBinaryHeap
	}{
		{
			name: "case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPQBasedOnBH(7, func(slice []interface{}, i, j int) bool {
				return slice[i].(int) < slice[j].(int)
			})
			pq.Push(7)
			assert.Equal(t, 7, pq.Peak().(int))
			pq.Push(6)
			assert.Equal(t, 6, pq.Peak().(int))
			pq.Push(8)
			assert.Equal(t, 6, pq.Peak().(int))
			pq.Push(3)
			assert.Equal(t, 3, pq.Peak().(int))
			assert.Equal(t, 3, pq.Pop().(int))
			assert.Equal(t, 3, pq.Size())
			assert.Equal(t, 6, pq.Pop().(int))
			assert.Equal(t, 2, pq.Size())
			assert.Equal(t, 7, pq.Pop().(int))
			assert.Equal(t, 1, pq.Size())
			assert.Equal(t, 8, pq.Pop().(int))
			assert.Equal(t, 0, pq.Size())
			assert.Equal(t, nil, pq.Pop())
			t.Log(pq.Peak())
		})
	}
}
