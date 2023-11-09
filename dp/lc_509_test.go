package dp

import (
	"testing"
)

func Test_fibdp(t *testing.T) {
	type args struct {
		n   int
		mem []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibdp(tt.args.n, tt.args.mem); got != tt.want {
				t.Errorf("fibdp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 0",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "n = 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "n = 2",
			args: args{
				n: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
