package backtrack

import (
	"reflect"
	"testing"
)

func Test_backtrack(t *testing.T) {
	type args struct {
		nums   []int
		track  []int
		used   map[int]struct{}
		answer *[][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "lc case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backtrack(tt.args.nums, tt.args.track, tt.args.used, tt.args.answer)
		})
	}
}

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		enabled bool
		args    args
		want    [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1},
			},
			want: [][]int{{1}},
		},
		{
			name:    "lc case 4",
			enabled: true,
			args: args{
				nums: []int{5, 4, 6, 2},
			},
			want: [][]int{{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.enabled {
				if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("permute() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_permutePV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		enabled bool
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1},
			},
			want: [][]int{{1}},
		},
		{
			name:    "lc case 4",
			enabled: true,
			args: args{
				nums: []int{5, 4, 6, 2},
			},
			want: [][]int{{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutePV1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutePV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permutePV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1},
			},
			want: [][]int{{1}},
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{5, 4, 6, 2},
			},
			want: [][]int{{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutePV2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutePV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
