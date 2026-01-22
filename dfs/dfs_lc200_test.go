package dfs

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				grid: [][]byte{
					{1, 1, 1, 1, 0},
					{1, 1, 0, 1, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: 1,
		},
		{
			name: "lc case 2",
			args: args{
				grid: [][]byte{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 1},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslandsPV1(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				grid: [][]byte{
					{1, 1, 1, 1, 0},
					{1, 1, 0, 1, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: 1,
		},
		{
			name: "lc case 2",
			args: args{
				grid: [][]byte{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 1},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslandsPV1(tt.args.grid); got != tt.want {
				t.Errorf("numIslandsPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslandsPV2(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				grid: [][]byte{
					{1, 1, 1, 1, 0},
					{1, 1, 0, 1, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: 1,
		},
		{
			name: "lc case 2",
			args: args{
				grid: [][]byte{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 1},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslandsPV2(tt.args.grid); got != tt.want {
				t.Errorf("numIslandsPV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslandsBFS(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				grid: [][]byte{
					{1, 1, 1, 1, 0},
					{1, 1, 0, 1, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
			want: 1,
		},
		{
			name: "lc case 2",
			args: args{
				grid: [][]byte{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 1},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.grid)
			if got := numIslandsBFS(tt.args.grid); got != tt.want {
				t.Errorf("numIslandsBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslandsV2(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 1,
		},
		{
			name: "lc case 2",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslandsV2(tt.args.grid); got != tt.want {
				t.Errorf("numIslandsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
