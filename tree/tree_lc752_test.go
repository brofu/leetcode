package tree

import (
	"testing"
)

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				deadends: []string{"0201", "0101", "0102", "1212", "2002"},
				target:   "0202",
			},
			want: 6,
		},
		{
			name: "lc case 2",
			args: args{
				deadends: []string{"8888"},
				target:   "0009",
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
				target:   "8888",
			},
			want: -1,
		},
		{
			name: "lc case 4",
			args: args{
				deadends: []string{"0000"},
				target:   "8888",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openLockV2(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				deadends: []string{"0201", "0101", "0102", "1212", "2002"},
				target:   "0202",
			},
			want: 6,
		},
		{
			name: "lc case 2",
			args: args{
				deadends: []string{"8888"},
				target:   "0009",
			},
			want: 1,
		},
		{
			name: "lc case 3",
			args: args{
				deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
				target:   "8888",
			},
			want: -1,
		},
		{
			name: "lc case 4",
			args: args{
				deadends: []string{"0000"},
				target:   "8888",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLockV2(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLockV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
