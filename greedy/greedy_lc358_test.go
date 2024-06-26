package greedy

import "testing"

func Test_rearrangeString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "lc case 1",
			args: args{
				s: "aabbcc",
				k: 3,
			},
			want: "abcabc",
		},
		{
			name: "lc case 2",
			args: args{
				s: "aaabc",
				k: 3,
			},
			want: "",
		},
		{
			name: "lc case 3",
			args: args{
				s: "aaadbbcc",
				k: 2,
			},
			//want: "abacabcd", another possible solution.
			want: "abcdabca",
		},
		{
			name: "lc case 4",
			args: args{
				s: "cihefgcceegaaebajgcfchhegfcehjbchfbjdjgdcedd",
				k: 6,
			},
			//want: "abacabcd", another possible solution.
			want: "ceghdfceghjacebdfgcehjabcedfghcejabdcefghicj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeString(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("rearrangeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
