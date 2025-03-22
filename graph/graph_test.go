package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDWG(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *DWG
	}{
		{
			name: "case 1",
			args: args{
				n: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewDWG(tt.args.n)

			got.AddEdge(0, 1, 2)
			assert.Equal(t, true, got.HasEdge(0, 1))
			assert.Equal(t, 2, got.Weight(0, 1))

			got.AddEdge(9, 10, 3)
			assert.Equal(t, true, got.HasEdge(9, 10))
			assert.Equal(t, 3, got.Weight(9, 10))

			got.DeleteEdge(9, 9)
			assert.Equal(t, true, got.HasEdge(0, 1))
			assert.Equal(t, true, got.HasEdge(9, 10))

			got.DeleteEdge(0, 1)
			assert.Equal(t, false, got.HasEdge(0, 1))
			assert.Equal(t, true, got.HasEdge(9, 10))
		})
	}
}
