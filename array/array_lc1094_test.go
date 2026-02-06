package array

import "testing"

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				trips: [][]int{
					{2, 1, 5}, {3, 3, 7},
				},
				capacity: 4,
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				trips: [][]int{
					{2, 1, 5}, {3, 3, 7},
				},
				capacity: 5,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				trips: [][]int{
					{2, 1, 5}, {3, 5, 7},
				},
				capacity: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("flag0", tt.args.trips)
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
