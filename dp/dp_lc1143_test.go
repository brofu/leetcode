package dp

import (
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "lc case 3",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequenceDPTable(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "lc case 3",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequenceDPTable(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceDPTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequenceDPTableWithSpaceCompress(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "lc case 3",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequenceDPTableWithSpaceCompress(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceDPTableWithSpaceCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
