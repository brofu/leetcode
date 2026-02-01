package stack

import (
	"testing"
)

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		enabled bool
		args    args
		want    int
	}{
		{
			name:    "case 1",
			enabled: true,
			args: args{
				nums: []int{
					2, 6, 4, 8, 10, 9, 15,
				},
			},
			want: 5,
		},
		{
			name:    "case 2",
			enabled: true,
			args: args{
				nums: []int{
					1, 2, 3, 4,
				},
			},
			want: 0,
		},
		{
			name:    "case 3",
			enabled: true,
			args: args{
				nums: []int{
					1,
				},
			},
			want: 0,
		},
		{
			name:    "case 4",
			enabled: true,
			args: args{
				nums: []int{
					4, 2, 3, 1,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		if !tt.enabled {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findUnsortedSubarrayV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		enabled bool
		args    args
		want    int
	}{
		{
			name:    "case 1",
			enabled: true,
			args: args{
				nums: []int{
					2, 6, 4, 8, 10, 9, 15,
				},
			},
			want: 5,
		},
		{
			name:    "case 2",
			enabled: true,
			args: args{
				nums: []int{
					1, 2, 3, 4,
				},
			},
			want: 0,
		},
		{
			name:    "case 3",
			enabled: true,
			args: args{
				nums: []int{
					1,
				},
			},
			want: 0,
		},
		{
			name:    "case 4",
			enabled: true,
			args: args{
				nums: []int{
					4, 2, 3, 1,
				},
			},
			want: 4,
		},
		{
			name:    "case 5",
			enabled: true,
			args: args{
				nums: []int{
					1, 2, 3, 3, 3,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarrayV2(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarrayV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
