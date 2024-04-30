package list

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head []int
		k    int
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
				k:    2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "lc case 2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := reverseKGroup(list, tt.args.k); !reflect.DeepEqual(got, want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroupV2(t *testing.T) {
	type args struct {
		head []int
		k    int
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
				k:    2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "lc case 2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := reverseKGroupV2(list, tt.args.k); !reflect.DeepEqual(got, want) {
				t.Errorf("case: %s, reverseKGroup() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_reverseKGroupP1(t *testing.T) {
	type args struct {
		head []int
		k    int
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
				k:    2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "lc case 2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
		{
			name: "self case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConstructListNodeFromSlice(tt.args.head)
			want := ConstructListNodeFromSlice(tt.want)
			if got := reverseKGroupP1(list, tt.args.k); !reflect.DeepEqual(got, want) {
				t.Errorf("case: %s, reverseKGroup() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
