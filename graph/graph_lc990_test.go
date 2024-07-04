package graph

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				equations: []string{"a==b", "b!=a"},
			},
			want: false,
		},
		{
			name: "lc case 2",
			args: args{
				equations: []string{"b==a", "a==b"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
