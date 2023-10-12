package list

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_partition(t *testing.T) {
	type args struct {
		list []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal",
			args: args{
				list: []int{1, 4, 3, 2, 5, 2},
				x:    3,
			},
			want: []int{1, 2, 2, 4, 3, 5},
		},
		{
			name: "empty",
			args: args{
				list: []int{},
				x:    3,
			},
			want: []int{},
		},
		{
			name: "one element",
			args: args{
				list: []int{1},
				x:    3,
			},
			want: []int{1},
		},
		{
			name: "two elements",
			args: args{
				list: []int{1, 3},
				x:    3,
			},
			want: []int{1, 3},
		},
		{
			name: "two elements ",
			args: args{
				list: []int{3, 5},
				x:    3,
			},
			want: []int{3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.list)
			want := ConstructListNodeFromSlice(tt.want)
			got := partition(input, tt.args.x)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSlitFromList(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{
				list: []int{},
			},
			want: []int{},
		},
		{
			name: "non empty",
			args: args{
				list: []int{3, 2, 5},
			},
			want: []int{3, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.list)
			if got := GetSliceFromList(input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSlitFromList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "lc 2",
			args: args{
				head: []int{1},
				n:    1,
			},
			want: []int{},
		},
		{
			name: "lc 3",
			args: args{
				head: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			got := removeNthFromEnd(input, tt.args.n)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, want)
				t.Error(tt.name, GetSliceFromList(got), tt.want)
			}
		})
	}
}

func TestFindFromEnd(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal",
			args: args{
				head: []int{1, 3, 2, 5},
				n:    1,
			},
			want: []int{5},
		},
		{
			name: "edge - one element",
			args: args{
				head: []int{5},
				n:    1,
			},
			want: []int{5},
		},
		{
			name: "edge - at the begining",
			args: args{
				head: []int{1, 3, 2, 5},
				n:    4,
			},
			want: []int{1, 3, 2, 5},
		},
		{
			name: "edge - at the ending",
			args: args{
				head: []int{1, 3, 2, 5},
				n:    1,
			},
			want: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := FindFromEnd(input, tt.args.n); !reflect.DeepEqual(got, want) {
				t.Errorf("Case: %s, FindFromEnd() = %v, want %v", tt.name, got, tt.want)
				t.Error(tt.name, GetSliceFromList(got), tt.want)
			}
		})
	}
}

func TestFindFromEndWithProtection(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal",
			args: args{
				head: []int{1, 3, 2, 5},
				n:    1,
			},
			want: []int{5},
		},
		{
			name: "edge - one element",
			args: args{
				head: []int{5},
				n:    1,
			},
			want: []int{5},
		},
		{
			name: "edge - empty",
			args: args{
				head: []int{},
				n:    1,
			},
			want: []int{},
		},
		{
			name: "edge - at the begining",
			args: args{
				head: []int{1, 3, 2, 5},
				n:    4,
			},
			want: []int{1, 3, 2, 5},
		},
		{
			name: "edge - at the ending",
			args: args{
				head: []int{1, 3, 2, 5},
				n:    1,
			},
			want: []int{5},
		},
		{
			name: "edge - n > length of list",
			args: args{
				head: []int{1, 3, 2, 5},
				n:    5,
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := FindFromEndWithProtection(input, tt.args.n); !reflect.DeepEqual(got, want) {
				t.Errorf("Case: %s, FindFromEnd() = %v, want %v", tt.name, got, tt.want)
				t.Error(tt.name, GetSliceFromList(got), tt.want)
			}
		})
	}
}

func Test_removeNthFromEndV2(t *testing.T) {

	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "lc 2",
			args: args{
				head: []int{1},
				n:    1,
			},
			want: []int{},
		},
		{
			name: "lc 3",
			args: args{
				head: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			got := removeNthFromEndV2(input, tt.args.n)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, want)
				t.Error(tt.name, GetSliceFromList(got), tt.want)
			}
		})
	}

}

func Test_removeNthFromEndV3(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "lc 2",
			args: args{
				head: []int{1},
				n:    1,
			},
			want: []int{},
		},
		{
			name: "lc 3",
			args: args{
				head: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			got := removeNthFromEndV3(input, tt.args.n)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, want)
				t.Error(tt.name, GetSliceFromList(got), tt.want)
			}
		})
	}
}

func Test_middleNode(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc-case-1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
			},
			want: []int{3, 4, 5},
		},
		{
			name: "lc-case-2",
			args: args{
				head: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := middleNode(input); !reflect.DeepEqual(got, want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc-case-1",
			args: args{
				head: []int{1, 3, 4, 7, 1, 2, 6},
			},
			want: []int{1, 3, 4, 1, 2, 6},
		},
		{
			name: "lc-case-1",
			args: args{
				head: []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "lc-case-3",
			args: args{
				head: []int{2, 1},
			},
			want: []int{2},
		},
		{
			name: "only one element",
			args: args{
				head: []int{2},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := deleteMiddleV2(input); !reflect.DeepEqual(got, want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
				t.Error(tt.name, GetSliceFromList(got), tt.want)
			}
		})
	}
}

func Test_deleteMiddleWithTwoLoop(t *testing.T) {

	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc-case-1",
			args: args{
				head: []int{1, 3, 4, 7, 1, 2, 6},
			},
			want: []int{1, 3, 4, 1, 2, 6},
		},
		{
			name: "lc-case-1",
			args: args{
				head: []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "lc-case-3",
			args: args{
				head: []int{2, 1},
			},
			want: []int{2},
		},
		{
			name: "only one element",
			args: args{
				head: []int{2},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := deleteMiddleWithTwoLoop(input); !reflect.DeepEqual(got, want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
				t.Error(tt.name, GetSliceFromList(got), tt.want)
			}
		})
	}

}

func Test_detectCycleV2(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc-case-1",
			args: args{
				head: []int{3, 2, 0, -4},
				n:    0,
			},
			want: []int{3},
		},
		{
			name: "lc-case-2",
			args: args{
				head: []int{1, 2},
				n:    0,
			},
			want: []int{1},
		},
		{
			name: "lc-case-3",
			args: args{
				head: []int{1},
				n:    -1,
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ConstructListNodeFromSliceWithCycle(tt.args.head, tt.args.n)
			got := detectCycleV2(input)
			if len(tt.want) == 0 {
				var empty *ListNode
				assert.Equal(t, empty, got)
			} else {
				assert.Equal(t, tt.want[0], got.Val)
			}
		})
	}
}
