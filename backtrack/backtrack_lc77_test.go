package backtrack

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combineV2(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineV2(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combineV3(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{{1}},
		},
		{
			name: "lc case 3",
			args: args{
				n: 5,
				k: 4,
			},
			want: [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineV3(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combineV4(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "lc case 1",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "lc case 2",
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{{1}},
		},
		{
			name: "lc case 3",
			args: args{
				n: 5,
				k: 4,
			},
			want: [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineV4(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
