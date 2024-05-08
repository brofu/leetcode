package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{5, 2, 3, 1},
				low:  0,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
				low:  0,
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.high = len(tt.args.nums) - 1
			QuickSort(tt.args.nums, tt.args.low, tt.args.high)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_quickSortPartition(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSortPartition(tt.args.nums, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("quickSortPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
