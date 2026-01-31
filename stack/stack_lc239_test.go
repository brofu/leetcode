package stack

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		enabled bool
		args    args
		want    []int
	}{
		{
			name:    "case 1",
			enabled: true,
			args: args{
				nums: []int{
					1, 3, -1, -3, 5, 3, 6, 7,
				},
				k: 3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:    "case 2",
			enabled: true,
			args: args{
				nums: []int{
					1,
				},
				k: 1,
			},
			want: []int{1},
		},
		{
			name:    "case 3",
			enabled: true,
			args: args{
				nums: []int{
					-7, -8, 7, 5, 7, 1, 6, 0,
				},
				k: 4,
			},
			want: []int{7, 7, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		if !tt.enabled {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlidingWindowV2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		enabled bool
		args    args
		want    []int
	}{
		{
			name:    "case 1",
			enabled: true,
			args: args{
				nums: []int{
					1, 3, -1, -3, 5, 3, 6, 7,
				},
				k: 3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:    "case 2",
			enabled: true,
			args: args{
				nums: []int{
					1,
				},
				k: 1,
			},
			want: []int{1},
		},
		{
			name:    "case 3",
			enabled: true,
			args: args{
				nums: []int{
					-7, -8, 7, 5, 7, 1, 6, 0,
				},
				k: 4,
			},
			want: []int{7, 7, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindowV2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindowV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
