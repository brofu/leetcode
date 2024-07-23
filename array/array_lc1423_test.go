package array

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				cardPoints: []int{1, 2, 3, 4, 5, 6, 1},
				k:          3,
			},
			want: 12,
		},
		{
			name: "lc case 2",
			args: args{
				cardPoints: []int{2, 2, 2},
				k:          2,
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				cardPoints: []int{9, 7, 7, 9, 7, 7, 9},
				k:          7,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
