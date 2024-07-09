package tree

import (
	"testing"
)

func Test_numTrees(t *testing.T) {
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
				n: 3,
			},
			want: 5,
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
			if got := numTrees(tt.args.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numTreesPV1(t *testing.T) {
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
				n: 3,
			},
			want: 5,
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
			if got := numTreesPV1(tt.args.n); got != tt.want {
				t.Errorf("numTreesPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
