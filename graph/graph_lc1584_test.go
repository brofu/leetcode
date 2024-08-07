package graph

import (
	"testing"
)

func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				points: [][]int{
					{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
				},
			},
			want: 20,
		},
		{
			name: "lc case 2",
			args: args{
				points: [][]int{
					{3, 12}, {-2, 5}, {-4, 1},
				},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostConnectPointsWithoutUF(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				points: [][]int{
					{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
				},
			},
			want: 20,
		},
		{
			name: "lc case 2",
			args: args{
				points: [][]int{
					{3, 12}, {-2, 5}, {-4, 1},
				},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPointsWithoutUF(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPointsWithoutUF() = %v, want %v", got, tt.want)
			}
		})
	}
}
