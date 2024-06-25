package dp

import (
	"testing"
)

func Test_convertArray(t *testing.T) {
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
				nums: []int{3, 2, 4, 5, 0},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{2, 2, 3, 4},
			},
			want: 0,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{15, 17, 8, 1, 2, 7, 5},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertArray(tt.args.nums); got != tt.want {
				t.Errorf("convertArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertArrayHeap(t *testing.T) {
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
				nums: []int{3, 2, 4, 5, 0},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{2, 2, 3, 4},
			},
			want: 0,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{15, 17, 8, 1, 2, 7, 5},
			},
			want: 11,
		},
		{
			name: "self case 1",
			args: args{
				nums: []int{10, 6, 5, 4, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertArrayHeap(tt.args.nums); got != tt.want {
				t.Errorf("convertArrayHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
