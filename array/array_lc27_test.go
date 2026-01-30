package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNuns []int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want:     2,
			wantNuns: []int{2, 2},
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want:     5,
			wantNuns: []int{0, 1, 3, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			assert.Equal(t, tt.want, got)
			for i := 0; i < got; i++ {
				assert.Equal(t, tt.wantNuns[i], tt.args.nums[i])
			}
		})
	}
}

func Test_removeElementPV1(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementPV1(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElementPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
