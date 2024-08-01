package array

import (
	"reflect"
	"testing"

	"github.com/brofu/leetcode/common"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head  *common.ListNode
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{1, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{1, 1, 2, 3, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "lc case 3",
			args: args{
				input: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.head = common.GenerateListNodeFromSlice(tt.args.input)
			got := deleteDuplicates(tt.args.head)
			gotSlice := common.GenerateSliceFromListNode(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
