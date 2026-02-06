package math

import (
	"testing"
)

func Test_superPow(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "lc case 2",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 1024,
		},
		{
			name: "lc case 3",
			args: args{
				a: 1,
				b: []int{4, 3, 3, 8, 5, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_superPowV2(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "lc case 2",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 1024,
		},
		{
			name: "lc case 3",
			args: args{
				a: 1,
				b: []int{4, 3, 3, 8, 5, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPowV2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPowV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_superPowV3(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "lc case 2",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 1024,
		},
		{
			name: "lc case 3",
			args: args{
				a: 1,
				b: []int{4, 3, 3, 8, 5, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPowV3(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPowV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_superPowV4(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "lc case 2",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 1024,
		},
		{
			name: "lc case 3",
			args: args{
				a: 1,
				b: []int{4, 3, 3, 8, 5, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPowV4(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPowV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_superPowV5(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "lc case 2",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 1024,
		},
		{
			name: "lc case 3",
			args: args{
				a: 1,
				b: []int{4, 3, 3, 8, 5, 2},
			},
			want: 1,
		},
		{
			name: "lc case 4",
			args: args{
				a: 2147483647,
				b: []int{2, 0, 0},
			},
			want: 1198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPowV5(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPowV5() = %v, want %v", got, tt.want)
			}
		})
	}
}
