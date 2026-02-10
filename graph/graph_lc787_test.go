package graph

import "testing"

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 4,
				flights: [][]int{
					{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200},
				},
				src: 0,
				dst: 3,
				k:   1,
			},
			want: 700,
		},
		{
			name: "case 2",
			args: args{
				n: 4,
				flights: [][]int{
					{0, 1, 100}, {1, 2, 100}, {0, 2, 500},
				},
				src: 0,
				dst: 2,
				k:   1,
			},
			want: 200,
		},
		{
			name: "case 3",
			args: args{
				n: 3,
				flights: [][]int{
					{0, 1, 100}, {1, 2, 100}, {0, 2, 500},
				},
				src: 0,
				dst: 2,
				k:   0,
			},
			want: 500,
		},
		{
			name: "case 3",
			args: args{
				n: 3,
				flights: [][]int{
					{0, 1, 100}, {1, 2, 100}, {0, 2, 500},
				},
				src: 0,
				dst: 2,
				k:   0,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("input", tt.args.flights)
			if got := findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k); got != tt.want {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
