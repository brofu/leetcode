package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal case 1",
			args: args{
				nums: []int{12, 11, 13, 5, 6, 7},
			},
			want: []int{5, 6, 7, 11, 12, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				panic("error")
			}
		})
	}
}

func TestHeapSortIterative(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal case 1",
			args: args{
				nums: []int{12, 11, 13, 5, 6, 7},
			},
			want: []int{5, 6, 7, 11, 12, 13},
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			HeapSortIterative(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("HeapSortIterative() = %v, want %v", tt.args.nums, tt.want)
			}

		})
	}
}
