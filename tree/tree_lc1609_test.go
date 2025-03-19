package tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isEvenOddTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 10, 4, 3, math.MaxInt, 7, 9, 12, 8, 6, math.MaxInt, math.MaxInt, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := GenerateTreeNodeFromSliceBFS(tt.args.nums)
			got := isEvenOddTree(root)
			assert.Equal(t, tt.want, got)
		})
	}
}
