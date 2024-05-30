package stack

import (
	"testing"
)

func Test_calculateWithParenthese(t *testing.T) {
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
			name: "lc case 1",
			args: args{
				s: "1+1",
			},
			want: 2,
		},

		{
			name: "lc case 2",
			args: args{
				s: "6-4/2",
			},
			want: 4,
		},

		{
			name: "self case 1",
			args: args{
				s: "2*(5+1)",
			},
			want:    12,
			enabled: true,
		},
		{
			name: "self case 2",
			args: args{
				s: "2*(5+5*2)",
			},
			want:    30,
			enabled: false,
		},

		{
			name: "lc case 3",
			args: args{
				s: "2*(5+5*2)/3+(6/2+8)",
			},
			want: 21,
		},
		{
			name: "lc case 4",
			args: args{
				s: "(4)+(2)",
			},
			want: 6,
		},
		{
			name: "lc case 5",
			args: args{
				s: "5+3-4-(1+2-7+(10-1+3+5+(3-0+(8-(3+(8-(10-(6-10-8-7+(0+0+7)-10+5-3-2+(9+0+(7+(2-(2-(9)-2+5+4+2+(2+9+1+5+5-8-9-2-9+1+0)-(5-(9)-(0-(7+9)+(10+(6-4+6))+0-2+(10+7+(8+(7-(8-(3)+(2)+(10-6+10-(2)-7-(2)+(3+(8))+(1-3-8)+6-(4+1)+(6))+6-(1)-(10+(4)+(8)+(5+(0))+(3-(6))-(9)-(4)+(2))))))-1)))+(9+6)+(0))))+3-(1))+(7))))))))",
			},
			want: -35,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWithParenthese(tt.args.s); got != tt.want {
				t.Errorf("calculateWithParenthese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateWithParentheseRecursively(t *testing.T) {
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
			name: "lc case 1",
			args: args{
				s: "1+1",
			},
			want: 2,
		},

		{
			name: "lc case 2",
			args: args{
				s: "6-4/2",
			},
			want: 4,
		},

		{
			name: "self case 1",
			args: args{
				s: "2*(5+1)",
			},
			want: 12,
		},
		{
			name: "self case 2",
			args: args{
				s: "2*(5+5*2)",
			},
			want:    30,
			enabled: false,
		},

		{
			name: "lc case 3",
			args: args{
				s: "2*(5+5*2)/3+(6/2+8)",
			},
			want:    21,
			enabled: true,
		},
		{
			name: "lc case 3-1",
			args: args{
				s: "2*(5+5*2)/3",
			},
			want:    10,
			enabled: true,
		},

		{
			name: "lc case 4",
			args: args{
				s: "(4)+(2)",
			},
			want: 6,
		},
		{
			name: "lc case 5",
			args: args{
				s: "5+3-4-(1+2-7+(10-1+3+5+(3-0+(8-(3+(8-(10-(6-10-8-7+(0+0+7)-10+5-3-2+(9+0+(7+(2-(2-(9)-2+5+4+2+(2+9+1+5+5-8-9-2-9+1+0)-(5-(9)-(0-(7+9)+(10+(6-4+6))+0-2+(10+7+(8+(7-(8-(3)+(2)+(10-6+10-(2)-7-(2)+(3+(8))+(1-3-8)+6-(4+1)+(6))+6-(1)-(10+(4)+(8)+(5+(0))+(3-(6))-(9)-(4)+(2))))))-1)))+(9+6)+(0))))+3-(1))+(7))))))))",
			},
			want: -35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if !tt.enabled {
			//return
			//}
			if got := calculateWithParentheseRecursively(tt.args.s); got != tt.want {
				t.Errorf("calculateWithParentheseRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}
