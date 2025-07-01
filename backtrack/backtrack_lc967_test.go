package backtrack

import (
	"reflect"
	"testing"
)

func Test_numsSameConsecDiff(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				n: 3,
				k: 7,
			},
			want: []int{181, 292, 707, 818, 929},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numsSameConsecDiff(tt.args.n, tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numsSameConsecDiff() = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}

func Test_numsSameConsecDiffV2(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				n: 3,
				k: 7,
			},
			want: []int{181, 292, 707, 818, 929},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numsSameConsecDiffV2(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numsSameConsecDiffV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
