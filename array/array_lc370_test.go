package array

import (
	"reflect"
	"testing"
)

func Test_getModifiedArray(t *testing.T) {
	type args struct {
		length  int
		updates [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				length: 5,
				updates: [][]int{
					{1, 3, 2}, {2, 4, 3}, {0, 2, -2},
				},
			},
			want: []int{-2, 0, 3, 5, 3},
		},
		{
			name: "case 2",
			args: args{
				length: 10,
				updates: [][]int{
					{2, 4, 6}, {5, 6, 8}, {1, 9, -4},
				},
			},
			want: []int{0, -4, 2, 2, 2, 4, 4, -4, -4, -4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getModifiedArray(tt.args.length, tt.args.updates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getModifiedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
