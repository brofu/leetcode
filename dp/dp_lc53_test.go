package dp

import (
	"testing"
)

func Test_maxSubArray(t *testing.T) {
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
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "lc case 5",
			args: args{
				nums: []int{-2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArrayPV1(t *testing.T) {
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
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "lc case 5",
			args: args{
				nums: []int{-2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayPV1(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArrayPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArrayPreSum(t *testing.T) {
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
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "lc case 5",
			args: args{
				nums: []int{-2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayPreSum(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArrayPreSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
