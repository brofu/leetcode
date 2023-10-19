package list

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "lc case 2",
			args: args{
				head: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			name: "lc case 3",
			args: args{
				head: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			got := reverseList(list)
			want := ConstructListNodeFromSlice(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListV2(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "lc case 2",
			args: args{
				head: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			name: "lc case 3",
			args: args{
				head: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			got := reverseListV2(list)
			want := ConstructListNodeFromSlice(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListV3(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "lc case 2",
			args: args{
				head: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			name: "lc case 3",
			args: args{
				head: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			got := reverseListV3(list)
			want := ConstructListNodeFromSlice(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListFirstN(t *testing.T) {
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
			name: "case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{2, 1, 3, 4, 5},
		},
		{
			name: "case 2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    5,
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "case 3",
			args: args{
				head: []int{1, 2},
				n:    0,
			},
			want: []int{1, 2},
		},
		{
			name: "case 4",
			args: args{
				head: []int{1, 2},
				n:    2,
			},
			want: []int{2, 1},
		},
		{
			name: "case 5",
			args: args{
				head: []int{},
				n:    2,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := reverseListFirstN(list, tt.args.n); !reflect.DeepEqual(got, want) {
				t.Errorf("reverseListFirstN() = %v, want %v", tt.args.head, tt.want)
			}
		})
	}
}

func Test_reverseListFirstNHelper(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name  string
		args  args
		want  *ListNode
		want1 *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := reverseListFirstNHelper(tt.args.head, tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListFirstNHelper() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("reverseListFirstNHelper() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_reverseListBetween(t *testing.T) {
	type args struct {
		head []int
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				m:    2,
				n:    4,
			},
			want: []int{1, 4, 3, 2, 5},
		},
		{
			name: "case 2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				m:    2,
				n:    5,
			},
			want: []int{1, 5, 4, 3, 2},
		},
		{
			name: "case 3",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				m:    3,
				n:    3,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "case 4",
			args: args{
				head: []int{1, 2},
				m:    3,
				n:    2,
			},
			want: []int{},
		},
		{
			name: "case 5",
			args: args{
				head: []int{1, 2},
				m:    3,
				n:    4,
			},
			want: []int{1, 2},
		},
		{
			name: "case 6",
			args: args{
				head: []int{1, 2, 3},
				m:    2,
				n:    4,
			},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := reverseListBetween(list, tt.args.m, tt.args.n); !reflect.DeepEqual(got, want) {
				t.Errorf("case: %s, reverseListBetween() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
