package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "one element",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},

		{
			name: "normal",
			args: args{
				nums: []int{3, 2, 1, 5},
			},
			want: []int{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.nums)
			//t.Log(tt.args.nums)
			//t.Log(tt.want)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.args.nums, tt.want)
			}

		})
	}
}
