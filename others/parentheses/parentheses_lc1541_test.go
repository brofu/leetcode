package parentheses

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
			name: "lc case 1",
			args: args{
				s: "(()))",
			},
			want: 1,
		},
		{
			name: "lc case 2",
			args: args{
				s: "())",
			},
			want: 0,
		},
		{
			name: "lc case 3",
			args: args{
				s: "))())(",
			},
			want: 3,
		},
		{
			name: "lc case 4",
			args: args{
				s: "(()))(()))()())))",
			},
			want: 4,
		},
		{
			name: "self case 1",
			args: args{
				s: ")())",
			},
			want: 2,
		},
		{
			name: "self case 2",
			args: args{
				s: "()()())))",
			},
			want: 3,
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
