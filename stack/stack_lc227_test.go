package stack

import (
	"testing"
)

func Test_calculateBasic(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s: "3+2*2",
			},
			want: 7,
		},

		{
			name: "lc case 2",
			args: args{
				s: " 3/2 ",
			},
			want: 1,
		},

		{
			name: "lc case 3",
			args: args{
				s: " 3+5 / 2 ",
			},
			want: 5,
		},
		{
			name: "lc case 4",
			args: args{
				s: "42",
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateBasic(tt.args.s); got != tt.want {
				t.Errorf("calculateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateBasicV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s: "3+2*2",
			},
			want: 7,
		},

		{
			name: "lc case 2",
			args: args{
				s: " 3/2 ",
			},
			want: 1,
		},

		{
			name: "lc case 3",
			args: args{
				s: " 3+5 / 2 ",
			},
			want: 5,
		},
		{
			name: "lc case 4",
			args: args{
				s: "42",
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateBasicV2(tt.args.s); got != tt.want {
				t.Errorf("calculateBasicV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
