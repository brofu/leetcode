package dijkstra

import "testing"

func Test_minimumEffortPath(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"lc case 1",
			args{
				[][]int{
					{1, 2, 2}, {3, 8, 2}, {5, 3, 5},
				},
			},
			2,
		},
		{
			"lc case 2",
			args{
				[][]int{
					{1, 2, 3}, {3, 8, 4}, {5, 3, 5},
				},
			},
			1,
		},
		{
			"lc case 3",
			args{
				[][]int{
					{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1},
				},
			},
			0,
		},
		{
			"lc case 4",
			args{
				[][]int{
					{10, 8}, {10, 8}, {1, 2}, {10, 3}, {1, 3}, {6, 3}, {5, 2},
				},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPath(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
