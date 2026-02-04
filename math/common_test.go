package math

import (
	"testing"
)

func TestQuickPow(t *testing.T) {
	type args struct {
		a int
		b int
		k int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				a: 2,
				b: 3,
				k: 5,
			},
			want: 3,
		},
		/*
			{
				name: "case 2",
				args: args{
					a: 2,
					b: -3,
					k: 5,
				},
				want: 3,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickPow(tt.args.a, tt.args.b, tt.args.k); got != tt.want {
				t.Errorf("QuickPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: 12,
				b: 8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCDV2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: 12,
				b: 8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCDV2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCDV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: 12,
				b: 8,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}
