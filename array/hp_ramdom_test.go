package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_triplev1(t *testing.T) {
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
			name: "case 1",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := triplev1(tt.args.nums, tt.args.target)
			got2 := triplev2(tt.args.nums, tt.args.target)
			t.Log(got1, got2)
			assert.Equal(t, got1, got2)
		})
	}
}
