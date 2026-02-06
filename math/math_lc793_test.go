package math

import "testing"

func Test_preimageSizeFZF(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				k: 0,
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				k: 5,
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				k: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preimageSizeFZF(tt.args.k); got != tt.want {
				t.Errorf("preimageSizeFZF() = %v, want %v", got, tt.want)
			}
		})
	}
}
