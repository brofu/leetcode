package math

import (
	"math"
	"testing"
)

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 3,
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				n: 5,
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				n: math.MaxInt64 >> 31,
			},
			want: 1073741816,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trailingZeroes(tt.args.n)
			t.Log("got", got)
			if got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
