package tree

import "testing"

func Test_countPalindromePaths(t *testing.T) {
	type args struct {
		parent []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"lc case 1",
			args{
				[]int{-1, 0, 0, 1, 1, 2},
				"acaabc",
			},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPalindromePaths(tt.args.parent, tt.args.s); got != tt.want {
				t.Errorf("countPalindromePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
