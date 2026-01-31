package stack

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				temperatures: []int{
					73, 74, 75, 71, 69, 72, 76, 73,
				},
			},
			want: []int{
				1, 1, 4, 2, 1, 1, 0, 0,
			},
		},
		{
			name: "case 2",
			args: args{
				temperatures: []int{
					30, 40, 50, 60,
				},
			},
			want: []int{
				1, 1, 1, 0,
			},
		},
		{
			name: "case 3",
			args: args{
				temperatures: []int{
					30, 60, 90,
				},
			},
			want: []int{
				1, 1, 0,
			},
		},
		{
			name: "case 4",
			args: args{
				temperatures: []int{
					89, 62, 70, 58, 47, 47, 46, 76, 100, 70,
				},
			},
			want: []int{
				8, 1, 5, 4, 3, 2, 1, 1, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dailyTemperaturesV2(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				temperatures: []int{
					73, 74, 75, 71, 69, 72, 76, 73,
				},
			},
			want: []int{
				1, 1, 4, 2, 1, 1, 0, 0,
			},
		},
		{
			name: "case 2",
			args: args{
				temperatures: []int{
					30, 40, 50, 60,
				},
			},
			want: []int{
				1, 1, 1, 0,
			},
		},
		{
			name: "case 3",
			args: args{
				temperatures: []int{
					30, 60, 90,
				},
			},
			want: []int{
				1, 1, 0,
			},
		},
		{
			name: "case 4",
			args: args{
				temperatures: []int{
					89, 62, 70, 58, 47, 47, 46, 76, 100, 70,
				},
			},
			want: []int{
				8, 1, 5, 4, 3, 2, 1, 1, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperaturesV2(tt.args.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperaturesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
