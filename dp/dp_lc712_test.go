package dp

import (
	"testing"
)

func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s1: "sea",
				s2: "eat",
			},
			want: 231,
		},
		{
			name: "lc case 2",
			args: args{
				s1: "delete",
				s2: "leet",
			},
			want: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumDeleteSumDPTable(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s1: "sea",
				s2: "eat",
			},
			want: 231,
		},
		{
			name: "lc case 2",
			args: args{
				s1: "delete",
				s2: "leet",
			},
			want: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSumDPTable(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSumDPTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumDeleteSumDPTableWithSpaceCompress(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s1: "sea",
				s2: "eat",
			},
			want: 231,
		},
		{
			name: "lc case 2",
			args: args{
				s1: "delete",
				s2: "leet",
			},
			want: 403,
		},
		{
			name: "lc case 3",
			args: args{
				s1: "a",
				s2: "at",
			},
			want: 116,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSumDPTableWithSpaceCompress(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSumDPTableWithSpaceCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
