package others

import (
	"testing"
)

func Test_resovle(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: []string{"abca", "ac", "abca", "abc", "abcc"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resovle(tt.args.s); got != tt.want {
				t.Errorf("resovle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resovleV2(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: []string{"abca", "ac", "abca", "abc", "abcc"},
			},
			want: "a",
		},
		{
			name: "case 2",
			args: args{
				s: []string{"abca", "abc", "abca", "abcd", "abcc"},
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resovleV2(tt.args.s); got != tt.want {
				t.Errorf("resovleV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
