package dp

import "testing"

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}},
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				envelopes: [][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
