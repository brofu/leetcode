package array

import (
	"testing"
)

func Test_quickSelectHelper(t *testing.T) {
	type args struct {
		array []int
		pivot int
		low   int
		high  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				array: []int{3, 2, 1, 5, 6, 4},
				pivot: 0,
				low:   0,
				high:  5,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				array: []int{2, 1, 3, 5, 6, 4},
				pivot: 4,
				low:   3,
				high:  5,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSelectHelper(tt.args.array, tt.args.pivot, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("quickSelectHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthLargestV2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargestV2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargestV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
