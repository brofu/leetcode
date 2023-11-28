package common

import (
	"reflect"
	"testing"
)

func Test_mergeSortWrapper(t *testing.T) {
	type args struct {
		nums  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSortWrapper(tt.args.nums, tt.args.start, tt.args.end)
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{8, 9, 2, 7, 10, 3},
			},
			want: []int{2, 3, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
