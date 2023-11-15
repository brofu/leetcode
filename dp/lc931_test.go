package dp

import "testing"

func Test_minFallingPathSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				matrix: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			},
			want: 13,
		},
		{
			name: "lc case 2",
			args: args{
				matrix: [][]int{{-19, 57}, {-40, -5}},
			},
			want: -59,
		},
		{
			name: "lc case 3",
			args: args{
				matrix: [][]int{{-51, -35, 74}, {-62, 14, -53}, {94, 61, -10}},
			},
			want: -98,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.args.matrix); got != tt.want {
				t.Errorf("CaseName: %s, minFallingPathSum() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
