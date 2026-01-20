package dc

import (
	"reflect"
	"sort"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		enabled bool
		args    args
		want    []int
	}{
		{
			name:    "case 1",
			enabled: true,
			args: args{
				expression: "2-1-1",
			},
			want: []int{0, 2},
		},
		{
			name:    "case 2",
			enabled: true,
			args: args{
				expression: "2*3-4*5",
			},
			want: []int{-34, -14, -10, -10, 10},
		},
		{
			name:    "case 3",
			enabled: true,
			args: args{
				expression: "11",
			},
			want: []int{11},
		}}
	for _, tt := range tests {
		if !tt.enabled {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			got := diffWaysToCompute(tt.args.expression)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
