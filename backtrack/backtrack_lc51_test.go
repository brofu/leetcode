package backtrack

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveNQueensPV1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueensPV1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueensPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveNQueensPV2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueensPV2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueensPV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveNQueensPV3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
			},
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
			},
			want: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueensPV3(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueensPV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
