package array

import (
	"reflect"
	"testing"
)

func Test_countOfAtoms(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				formula: "Mg(OH)2",
			},
			want: "H2MgO2",
		},
		{
			name: "case 2",
			args: args{
				formula: "K4(ON(SO3)2)2",
			},
			want: "K4N2O14S4",
		},
		{
			name: "case 3",
			args: args{
				formula: "H50",
			},
			want: "H50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("input", tt.args.formula)
			if got := countOfAtoms(tt.args.formula); got != tt.want {
				t.Errorf("countOfAtoms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOfAtomsHelper(t *testing.T) {
	type args struct {
		formula string
		start   int
		end     int
		pidx    map[int]int
	}
	tests := []struct {
		name string
		args args
		want map[byte]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfAtomsHelper(tt.args.formula, tt.args.start, tt.args.end, tt.args.pidx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countOfAtomsHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
