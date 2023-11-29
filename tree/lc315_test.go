package tree

import (
	"reflect"
	"testing"
)

func Test_countSmaller(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{5, 2, 6, 1},
			},
			want: []int{2, 1, 1, 0},
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{-1},
			},
			want: []int{0},
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{-1, -1},
			},
			want: []int{0, 0},
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
