package array

import "testing"

func Test_shortestPalindrome(t *testing.T) {
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
				s: "aacecaaa",
			},
			want: "aaacecaaa",
		},
		{
			name: "lc case 2",
			args: args{
				s: "abcd",
			},
			want: "dcbabcd",
		},
		{
			name: "lc case 3",
			args: args{
				s: "abb",
			},
			want: "bbabb",
		},
		{
			name: "lc case 4",
			args: args{
				s: "aabba",
			},
			want: "abbaabba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
