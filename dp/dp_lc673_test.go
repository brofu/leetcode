package dp

import "testing"

func Test_findNumberOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "self case 1",
			args: args{
				nums: []int{0, 1, 3, 2, 4},
			},
			want: 2,
		},
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 3, 5, 4, 7},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
