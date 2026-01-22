package dp

import (
	"testing"
)

func Test_minFallingPathSumDPA(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{2, 1, 3}, {6, 5, 4}, {7, 8, 9},
				},
			},
			want: 13,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{-19, 57}, {-40, -5},
				},
			},
			want: -59,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{17, 82}, {1, -44},
				},
			},
			want: -27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.matrix)
			if got := minFallingPathSumDPA(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSumDPA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minFallingPathSumDPF(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{2, 1, 3}, {6, 5, 4}, {7, 8, 9},
				},
			},
			want: 13,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{-19, 57}, {-40, -5},
				},
			},
			want: -59,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{17, 82}, {1, -44},
				},
			},
			want: -27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSumDPF(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSumDPF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minFallingPathSumDFS(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{2, 1, 3}, {6, 5, 4}, {7, 8, 9},
				},
			},
			want: 13,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{-19, 57}, {-40, -5},
				},
			},
			want: -59,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{17, 82}, {1, -44},
				},
			},
			want: -27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSumDFS(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSumDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minFallingPathSumV1(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{2, 1, 3}, {6, 5, 4}, {7, 8, 9},
				},
			},
			want: 13,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{-19, 57}, {-40, -5},
				},
			},
			want: -59,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{17, 82}, {1, -44},
				},
			},
			want: -27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSumV1(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSumV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minFallingPathSumV1WithMemo(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{2, 1, 3}, {6, 5, 4}, {7, 8, 9},
				},
			},
			want: 13,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{-19, 57}, {-40, -5},
				},
			},
			want: -59,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{17, 82}, {1, -44},
				},
			},
			want: -27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSumV1WithMemo(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSumV1WithMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minFallingPathSumV1WithDP(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{2, 1, 3}, {6, 5, 4}, {7, 8, 9},
				},
			},
			want: 13,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{-19, 57}, {-40, -5},
				},
			},
			want: -59,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{17, 82}, {1, -44},
				},
			},
			want: -27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSumV1WithDP(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSumV1WithDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
