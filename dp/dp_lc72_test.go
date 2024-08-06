package dp

import (
	"testing"
)

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistanceDPTable(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceDPTable(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistanceDPTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
