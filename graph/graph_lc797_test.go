package graph

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				graph: [][]int{
					{1, 2}, {3}, {3}, {},
				},
			},
			want: [][]int{
				{0, 1, 3}, {0, 2, 3},
			},
		},
		{
			name: "lc case 2",
			args: args{
				graph: [][]int{
					{4, 3, 1}, {3, 2, 4}, {3}, {4}, {},
				},
			},
			want: [][]int{
				{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
