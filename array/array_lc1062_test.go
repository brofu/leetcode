package array

import (
	"testing"
)

func Test_longestRepeatingSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s: "abcd",
			},
			want: 0,
		},
		{
			name: "lc case 2",
			args: args{
				s: "abbaba",
			},
			want: 2,
		},
		{
			name: "lc case 3",
			args: args{
				s: "aabcaabdaab",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestRepeatingSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestRepeatingSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestRepeatingSubstringV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s: "abcd",
			},
			want: 0,
		},
		{
			name: "lc case 2",
			args: args{
				s: "abbaba",
			},
			want: 2,
		},
		{
			name: "lc case 3",
			args: args{
				s: "aabcaabdaab",
			},
			want: 3,
		},
		{
			name: "lc case 4",
			args: args{
				s: "aaaaa",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestRepeatingSubstringV2(tt.args.s); got != tt.want {
				t.Errorf("longestRepeatingSubstringV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
