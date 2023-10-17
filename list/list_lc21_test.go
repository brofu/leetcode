package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc-case-1",
			args: args{
				head: []int{3, 2, 0, -4},
				n:    0,
			},
			want: true,
		},
		{
			name: "lc-case-2",
			args: args{
				head: []int{1, 2},
				n:    0,
			},
			want: true,
		},
		{
			name: "lc-case-3",
			args: args{
				head: []int{1},
				n:    -1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSliceWithCycle(tt.args.head, tt.args.n)
			got := hasCycle(input)
			assert.Equal(t, tt.want, got)
		})
	}
}
