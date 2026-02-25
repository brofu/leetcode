package dp

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		arr        []int
		difference int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arr:        []int{1, 2, 3, 4},
				difference: 1,
			},
			want: 4,
		},
		{
			name: "case 1",
			args: args{
				arr:        []int{1, 2, 3, 4},
				difference: 1,
			},
			want: 4,
		},
		{
			name: "case 1",
			args: args{
				arr:        []int{1, 2, 3, 4},
				difference: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.arr, tt.args.difference); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
