package sort

import "testing"

func Test_countRangeSum(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums:  []int{-2, 5, -1},
				lower: -2,
				upper: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRangeSum(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countRangeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countRangeSumWrapper(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRangeSumWrapper(tt.args.nums, tt.args.left, tt.args.right, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countRangeSumWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
