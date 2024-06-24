package backtrack

import "testing"

func Test_minTransfers(t *testing.T) {
	type args struct {
		transactions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				transactions: [][]int{{0, 1, 10}, {2, 0, 5}},
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				transactions: [][]int{{0, 1, 10}, {1, 0, 1}, {1, 2, 5}, {2, 0, 5}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTransfers(tt.args.transactions); got != tt.want {
				t.Errorf("minTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}
