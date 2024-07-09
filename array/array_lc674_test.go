package array

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 3, 5, 4, 7},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1, 3, 5, 7},
			},
			want: 4,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{1, 3, 5, 4, 2, 3, 4, 5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
