package array

import "testing"

func Test_longestWPI(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				hours: []int{9, 9, 6, 0, 6, 6, 9},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				hours: []int{6, 6, 6},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWPI(tt.args.hours); got != tt.want {
				t.Errorf("longestWPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
