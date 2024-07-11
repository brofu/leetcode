package array

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
