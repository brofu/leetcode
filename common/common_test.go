package common

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{8, 9, 2, 7, 10, 3},
			},
			want: []int{2, 3, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func TestMergeSortPV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{8, 9, 2, 7, 10, 3},
			},
			want: []int{2, 3, 7, 8, 9, 10},
		},
		{
			name: "normal case 1",
			args: args{
				nums: []int{12, 11, 13, 5, 6, 7},
			},
			want: []int{5, 6, 7, 11, 12, 13},
		},
		{
			name: "normal case 2",
			args: args{
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "normal case 3",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSortPV1(tt.args.nums)
		})
	}
}

func TestCheckIntSlice(t *testing.T) {
	type args struct {
		a []int
		b [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				a: []int{1, 2},
				b: [][]int{{1, 2, 3}, {1, 2}},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				a: []int{1, 2},
				b: [][]int{{1, 2, 3}, {1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIntSlice(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CheckIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeepEqualIntSlice(t *testing.T) {
	type args struct {
		a [][]int
		b [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				a: [][]int{{1, 2}},
				b: [][]int{{1, 2, 3}, {1, 2}},
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				a: [][]int{{1, 2, 3}, {1, 2}},
				b: [][]int{{1, 2, 3}, {1, 2}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeepEqualIntSlice(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("DeepEqualIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinString(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				a: "ab",
				b: "abc",
			},
			want: "ab",
		},
		{
			name: "case 2",
			args: args{
				a: "ba",
				b: "abc",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinString(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinString() = %v, want %v", got, tt.want)
			}
		})
	}
}
