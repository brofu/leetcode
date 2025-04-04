package graph

import (
	"testing"
)

func Test_possibleBipartition(t *testing.T) {
	type args struct {
		n        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
				dislikes: [][]int{
					{1, 2}, {1, 3}, {2, 4},
				},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				n: 3,
				dislikes: [][]int{
					{1, 2}, {1, 3}, {2, 3},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.args.n, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possibleBipartitionDFSPV1(t *testing.T) {
	type args struct {
		n        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
				dislikes: [][]int{
					{1, 2}, {1, 3}, {2, 4},
				},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				n: 3,
				dislikes: [][]int{
					{1, 2}, {1, 3}, {2, 3},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartitionDFSPV1(tt.args.n, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartitionDFSPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possibleBipartitionBFSPV1(t *testing.T) {
	type args struct {
		n        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
				dislikes: [][]int{
					{1, 2}, {1, 3}, {2, 4},
				},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				n: 3,
				dislikes: [][]int{
					{1, 2}, {1, 3}, {2, 3},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartitionBFSPV1(tt.args.n, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartitionDFSPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
