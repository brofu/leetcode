package backtrack

import (
	"testing"
)

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalNQueensPV1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueensPV1(tt.args.n); got != tt.want {
				t.Errorf("totalNQueensPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalNQueensPV2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueensPV2(tt.args.n); got != tt.want {
				t.Errorf("totalNQueensPV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalNQueensPV3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueensPV3(tt.args.n); got != tt.want {
				t.Errorf("totalNQueensPV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalNQueensPV4(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueensPV4(tt.args.n); got != tt.want {
				t.Errorf("totalNQueensPV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
