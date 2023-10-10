package list

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "both empty",
			args: args{
				list1: []int{},
				list2: []int{},
			},
			want: []int{},
		},
		{
			name: "one empty",
			args: args{
				list1: []int{1},
				list2: []int{},
			},
			want: []int{1},
		},

		{
			name: "one empty - v2",
			args: args{
				list1: []int{},
				list2: []int{2},
			},
			want: []int{2},
		},
		{
			name: "none empty",
			args: args{
				list1: []int{1},
				list2: []int{2},
			},
			want: []int{1, 2},
		},
		{
			name: "none empty",
			args: args{
				list1: []int{1},
				list2: []int{-1, 2},
			},
			want: []int{-1, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := ConstructListNodeFromSlice(tt.args.list1)
			list2 := ConstructListNodeFromSlice(tt.args.list2)
			want := ConstructListNodeFromSlice(tt.want)
			if got := MergeTwoLists(list1, list2); !reflect.DeepEqual(got, want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
