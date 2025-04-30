package dijkstra

import "testing"

func Test_maxProbability(t *testing.T) {
	type args struct {
		n          int
		edges      [][]int
		succProb   []float64
		start_node int
		end_node   int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"lc case 1",
			args{
				3,
				[][]int{
					{0, 1}, {1, 2}, {0, 2},
				},
				[]float64{0.5, 0.5, 0.2},
				0,
				2,
			},
			0.25,
		},
		{
			"lc case 2",
			args{
				3,
				[][]int{
					{0, 1}, {1, 2}, {0, 2},
				},
				[]float64{0.5, 0.5, 0.3},
				0,
				2,
			},
			0.3,
		},
		{
			"lc case 3",
			args{
				3,
				[][]int{
					{0, 1},
				},
				[]float64{0.5},
				0,
				2,
			},
			0,
		},
		{
			"lc case 4",
			args{
				5,
				[][]int{
					{1, 4}, {2, 4}, {0, 4}, {0, 3}, {0, 2}, {2, 3},
				},
				[]float64{0.37, 0.17, 0.93, 0.23, 0.39, 0.04},
				3,
				4,
			},
			0.21390,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProbability(tt.args.n, tt.args.edges, tt.args.succProb, tt.args.start_node, tt.args.end_node); got != tt.want {
				t.Errorf("maxProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
