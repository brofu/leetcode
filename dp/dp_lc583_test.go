package dp

import (
	"testing"
)

func Test_minDistance583(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				word1: "sea",
				word2: "eat",
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "leetcode",
				word2: "etco",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance583(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance583() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistance583WithSpaceCompress(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				word1: "sea",
				word2: "eat",
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "leetcode",
				word2: "etco",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance583WithSpaceCompress(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance583WithSpaceCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
