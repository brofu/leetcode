package array

import (
	"reflect"
	"testing"
)

func Test_pivotArray(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				nums:  []int{9, 12, 5, 10, 14, 3, 10},
				pivot: 10,
			},
			want: []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			name: "lc case 2",
			args: args{
				nums:  []int{-3, 4, 3, 2},
				pivot: 2,
			},
			want: []int{-3, 2, 4, 3},
		},
		{
			name: "empty",
			args: args{
				nums:  []int{},
				pivot: 10,
			},
			want: []int{},
		},
		{
			name: "one element, smaller",
			args: args{
				nums:  []int{9},
				pivot: 10,
			},
			want: []int{9},
		},
		{
			name: "one emlement, equal",
			args: args{
				nums:  []int{10},
				pivot: 10,
			},
			want: []int{10},
		},
		{
			name: "one element, larger",
			args: args{
				nums:  []int{9},
				pivot: 1,
			},
			want: []int{9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotArray(tt.args.nums, tt.args.pivot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pivotArrayV2(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				nums:  []int{9, 12, 5, 10, 14, 3, 10},
				pivot: 10,
			},
			want: []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			name: "lc case 2",
			args: args{
				nums:  []int{-3, 4, 3, 2},
				pivot: 2,
			},
			want: []int{-3, 2, 4, 3},
		},
		{
			name: "empty",
			args: args{
				nums:  []int{},
				pivot: 10,
			},
			want: []int{},
		},
		{
			name: "one element, smaller",
			args: args{
				nums:  []int{9},
				pivot: 10,
			},
			want: []int{9},
		},
		{
			name: "one emlement, equal",
			args: args{
				nums:  []int{10},
				pivot: 10,
			},
			want: []int{10},
		},
		{
			name: "one element, larger",
			args: args{
				nums:  []int{9},
				pivot: 1,
			},
			want: []int{9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotArrayV2(tt.args.nums, tt.args.pivot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArrayV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pivotArrayV3(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				nums:  []int{9, 12, 5, 10, 14, 3, 10},
				pivot: 10,
			},
			want: []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			name: "lc case 2",
			args: args{
				nums:  []int{-3, 4, 3, 2},
				pivot: 2,
			},
			want: []int{-3, 2, 4, 3},
		},
		{
			name: "empty",
			args: args{
				nums:  []int{},
				pivot: 10,
			},
			want: []int{},
		},
		{
			name: "one element, smaller",
			args: args{
				nums:  []int{9},
				pivot: 10,
			},
			want: []int{9},
		},
		{
			name: "one emlement, equal",
			args: args{
				nums:  []int{10},
				pivot: 10,
			},
			want: []int{10},
		},
		{
			name: "one element, larger",
			args: args{
				nums:  []int{9},
				pivot: 1,
			},
			want: []int{9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotArrayV3(tt.args.nums, tt.args.pivot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArrayV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
