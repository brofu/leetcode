package mst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCostConnectPointsPrim(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1 - side missing without full sides",
			args: args{
				points: [][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostConnectPointsPrim(tt.args.points)
			assert.Equal(t, tt.want, got)
		})
	}
}
