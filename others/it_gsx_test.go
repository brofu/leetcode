package others

import (
	"reflect"
	"testing"
)

func Test_resolve0730(t *testing.T) {
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
				nums: []int{3, 2, 1, 1},
			},
			want: []int{1, 1, -1, -1},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 2, 5, 1},
			},
			want: []int{-1, -1, -1, -1},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 1, 2, 3, 4, 2, 1},
			},
			want: []int{-1, -1, 1, 2, 2, 1, -1},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolve0730(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolve0730Stack(t *testing.T) {
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
				nums: []int{3, 2, 1, 1},
			},
			want: []int{1, 1, -1, -1},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 2, 5, 1},
			},
			want: []int{-1, -1, -1, -1},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 1, 2, 3, 4, 2, 1},
			},
			want: []int{-1, -1, 1, 2, 2, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolve0730Stack(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolve0730Stack() = %v, want %v", got, tt.want)
			}
		})
	}
}
