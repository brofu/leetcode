package array

import (
	"testing"
)

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{23, 2, 4, 6, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    13,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSubarraySumV2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{23, 2, 4, 6, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    13,
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{23, 2, 4, 6, 6},
				k:    7,
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{0, 0},
				k:    1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySumV2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
