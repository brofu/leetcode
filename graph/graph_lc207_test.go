package graph

import (
	"testing"
)

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canFinishPV1(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinishPV1(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinishPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canFinishBFS(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinishBFS(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinishBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
