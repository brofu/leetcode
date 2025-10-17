package dp

import "testing"

func Test_minInsertions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "zzazz",
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				s: "mbadm",
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				s: "leetcode",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minInsertionsDPT(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertionsDPT(tt.args.s); got != tt.want {
				t.Errorf("minInsertionsDPT() = %v, want %v", got, tt.want)
			}
		})
	}
}
