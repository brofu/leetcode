package array

import (
	"reflect"
	"testing"
)

func Test_arrayRankTransform(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				arr: []int{100, 100, 100},
			},
			want: []int{1, 1, 1},
		},
		{
			name: "case 1",
			args: args{
				arr: []int{40, 10, 20, 30},
			},
			want: []int{4, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayRankTransform(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
