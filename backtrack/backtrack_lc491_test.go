package backtrack

import (
	"reflect"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{4, 6, 7, 7},
			},
			want: [][]int{
				{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubsequences(tt.args.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
			t.Log(tt.want)
			t.Log(got)
		})
	}
}
