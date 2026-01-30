package array

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "lc case 1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "lc case 2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindromePV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "lc case 1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "lc case 2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromePV2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromePV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
