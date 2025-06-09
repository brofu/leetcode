package dfs

import (
	"testing"
)

func Test_closedIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 1, 1, 1, 0},
					{1, 0, 0, 0, 0, 1, 1, 0},
					{1, 0, 1, 0, 1, 1, 1, 0},
					{1, 0, 0, 0, 0, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 0},
				},
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				grid: [][]int{{0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0}},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				grid: [][]int{{1, 1, 1, 1, 1, 1, 1},
					{1, 0, 0, 0, 0, 0, 1},
					{1, 0, 1, 1, 1, 0, 1},
					{1, 0, 1, 0, 1, 0, 1},
					{1, 0, 1, 1, 1, 0, 1},
					{1, 0, 0, 0, 0, 0, 1},
					{1, 1, 1, 1, 1, 1, 1},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closedIsland(tt.args.grid); got != tt.want {
				t.Errorf("closedIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closedIslandPV1(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 1, 1, 1, 0},
					{1, 0, 0, 0, 0, 1, 1, 0},
					{1, 0, 1, 0, 1, 1, 1, 0},
					{1, 0, 0, 0, 0, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 0},
				},
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				grid: [][]int{{0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0}},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				grid: [][]int{{1, 1, 1, 1, 1, 1, 1},
					{1, 0, 0, 0, 0, 0, 1},
					{1, 0, 1, 1, 1, 0, 1},
					{1, 0, 1, 0, 1, 0, 1},
					{1, 0, 1, 1, 1, 0, 1},
					{1, 0, 0, 0, 0, 0, 1},
					{1, 1, 1, 1, 1, 1, 1},
				},
			},
			want: 2,
		},
		{
			name: "lc case 4",
			args: args{
				grid: [][]int{
					{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
					{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
					{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
					{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
					{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
					{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
					{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
					{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
					{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, r := range tt.args.grid {
				t.Log(r)
			}
			if got := closedIslandPV1(tt.args.grid); got != tt.want {
				t.Errorf("closedIslandPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closedIslandBFS(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 1, 1, 1, 0},
					{1, 0, 0, 0, 0, 1, 1, 0},
					{1, 0, 1, 0, 1, 1, 1, 0},
					{1, 0, 0, 0, 0, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 0},
				},
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				grid: [][]int{{0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0}},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				grid: [][]int{{1, 1, 1, 1, 1, 1, 1},
					{1, 0, 0, 0, 0, 0, 1},
					{1, 0, 1, 1, 1, 0, 1},
					{1, 0, 1, 0, 1, 0, 1},
					{1, 0, 1, 1, 1, 0, 1},
					{1, 0, 0, 0, 0, 0, 1},
					{1, 1, 1, 1, 1, 1, 1},
				},
			},
			want: 2,
		},
		{
			name: "lc case 4",
			args: args{
				grid: [][]int{
					{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
					{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
					{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
					{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
					{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
					{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
					{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
					{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
					{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closedIslandBFS(tt.args.grid); got != tt.want {
				t.Errorf("closedIslandBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
