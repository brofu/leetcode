package dfs

import "testing"

func Test_uniquePathsIII(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{
					{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1},
				},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{
					{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2},
				},
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				grid: [][]int{
					{0, 1}, {2, 0},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsIII(tt.args.grid); got != tt.want {
				t.Errorf("uniquePathsIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
