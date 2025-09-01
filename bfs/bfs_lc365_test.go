package bfs

import "testing"

func Test_canMeasureWater(t *testing.T) {
	type args struct {
		x      int
		y      int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				x:      3,
				y:      5,
				target: 4,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				x:      2,
				y:      6,
				target: 5,
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				x:      1,
				y:      2,
				target: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWater(tt.args.x, tt.args.y, tt.args.target); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
