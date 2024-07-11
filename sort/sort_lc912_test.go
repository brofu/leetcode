package sort

import (
	"reflect"
	"testing"
)

func Test_sortArrayPV1(t *testing.T) {
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
		{
			name: "normal case 1",
			args: args{
				nums: []int{12, 11, 13, 5, 6, 7},
			},
			want: []int{5, 6, 7, 11, 12, 13},
		},
		{
			name: "normal case 2",
			args: args{
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "normal case 3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayPV1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
