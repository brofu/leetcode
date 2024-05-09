package examples

import "testing"

func TestCalculateExpression(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "only add",
			args: args{
				expr: "1+1",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateExpression(tt.args.expr); got != tt.want {
				t.Errorf("CalculateExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
