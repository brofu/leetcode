package list

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				head: []int{1, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "case 2",
			args: args{
				head: []int{1, 1, 2, 3, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			got := deleteDuplicates(list)
			gotList := GetSliceFromList(got)
			if !reflect.DeepEqual(gotList, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", gotList, tt.want)
			}
		})
	}
}
