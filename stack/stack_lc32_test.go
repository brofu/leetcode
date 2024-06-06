package stack

import (
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		enabled bool
	}{
		{
			name: "lc ase 1",
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: "lc ase 2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "lc ase 3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "lc ase 4",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.enabled {
				return
			}
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParenthesesV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc ase 1",
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: "lc ase 2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "lc ase 3",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesV2(tt.args.s); got != tt.want {
				t.Errorf("longestValidParenthesesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParenthesesV3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc ase 1",
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: "lc ase 2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "lc ase 3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "lc ase 4",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesV3(tt.args.s); got != tt.want {
				t.Errorf("longestValidParenthesesV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParenthesesV4(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc ase 1",
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: "lc ase 2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "lc ase 3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "lc ase 4",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesV4(tt.args.s); got != tt.want {
				t.Errorf("longestValidParenthesesV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParenthesesV5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc ase 1",
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: "lc ase 2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "lc ase 3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "lc ase 4",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesV5(tt.args.s); got != tt.want {
				t.Errorf("longestValidParenthesesV5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParenthesesV3PV1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc ase 1",
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: "lc ase 2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "lc ase 3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "lc ase 4",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesV3PV1(tt.args.s); got != tt.want {
				t.Errorf("longestValidParenthesesV3PV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
