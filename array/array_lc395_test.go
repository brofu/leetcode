package array

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s: "aaabb",
				k: 3,
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				s: "ababbc",
				k: 2,
			},
			want: 5,
		},
		{
			name: "lc case 3",
			args: args{
				s: "baaabcb",
				k: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
