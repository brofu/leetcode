package mst

import "testing"

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
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
					{2, 1, 1}, {2, 3, 1}, {3, 4, 1},
				},
				4,
				2,
			},
			2,
		},
		{
			"lc case 2",
			args{
				[][]int{
					{1, 2, 1},
				},
				2,
				1,
			},
			1,
		},
		{
			"lc case 3",
			args{
				[][]int{
					{1, 2, 1},
				},
				2,
				2,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime(tt.args.times, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
