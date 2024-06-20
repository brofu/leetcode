package tree

import "testing"

func Test_pathSumIV(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{113, 215, 221},
			},
			want: 12,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{113, 221},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSumIV(tt.args.nums); got != tt.want {
				t.Errorf("pathSumIV() = %v, want %v", got, tt.want)
			}
		})
	}
}
