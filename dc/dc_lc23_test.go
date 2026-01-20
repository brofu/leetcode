package dc

import (
	"reflect"
	"testing"

	"github.com/brofu/leetcode/common"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				input: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name: "lc case 2",
			args: args{
				input: [][]int{},
			},
			want: []int{},
		},
		{
			name: "lc case 3",
			args: args{
				input: [][]int{{}},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, input := range tt.args.input {
				tt.args.lists = append(tt.args.lists, common.GenerateListNodeFromSlice(input))
			}
			want := common.GenerateListNodeFromSlice(tt.want)
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, want) {
				t.Errorf("mergeKLists() = %v, want %v", common.GenerateSliceFromListNode(got), tt.want)
			}
		})
	}
}

func Test_mergeKListsRecursive(t *testing.T) {
	type args struct {
		list  []*ListNode
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKListsRecursive(tt.args.list, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKListsRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
