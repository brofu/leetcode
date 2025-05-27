package array

import (
	"testing"
)

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_characterReplacementV2(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacementV2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacementV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_characterReplacementPV3(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacementPV3(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacementV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
