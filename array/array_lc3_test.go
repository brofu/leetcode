package array

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "lc case2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "lc case3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "lc case2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "lc case3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "lc case4",
			args: args{
				s: "abba",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringV2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringPV1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "lc case2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "lc case3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "lc case4",
			args: args{
				s: " ",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringPV1(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
