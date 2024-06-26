package backtrack

import "testing"

func Test_threeSumSmaller(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums:   []int{-2, 0, 1, 3},
				target: 2,
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: 0,
		},
		{
			name: "lc case 1",
			args: args{
				nums:   []int{0},
				target: 0,
			},
			want: 0,
		},
		{
			name: "lc case 4",
			args: args{
				nums:   []int{1, 1, -2},
				target: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumSmaller(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
