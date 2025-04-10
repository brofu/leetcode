package mst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumCostPrim(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal case ",
			args: args{
				n:           3,
				connections: [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}},
			},
			want: 6,
		},
		{
			name: "normal case 2",
			args: args{
				n:           4,
				connections: [][]int{{1, 2, 3}, {3, 4, 4}},
			},
			want: -1,
		},
		{
			name: "normal case 3",
			args: args{
				n:           5,
				connections: [][]int{{2, 1, 50459}, {3, 2, 47477}, {4, 2, 52585}, {5, 3, 16477}},
			},
			want: 166998,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCostPrim(tt.args.n, tt.args.connections)
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_minimumCostPrimV2(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal case ",
			args: args{
				n:           3,
				connections: [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}},
			},
			want: 6,
		},
		{
			name: "normal case 2",
			args: args{
				n:           4,
				connections: [][]int{{1, 2, 3}, {3, 4, 4}},
			},
			want: -1,
		},
		{
			name: "normal case 3",
			args: args{
				n:           5,
				connections: [][]int{{2, 1, 50459}, {3, 2, 47477}, {4, 2, 52585}, {5, 3, 16477}},
			},
			want: 166998,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCostPrimV2(tt.args.n, tt.args.connections)
			assert.Equal(t, got, tt.want)
		})
	}
}
