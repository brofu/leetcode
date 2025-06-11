package graph

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: []int{0, 1},
		},
		{
			name: "lc case 2",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0}, {2, 0}, {3, 1}, {3, 2},
				},
			},
			want: []int{0, 2, 1, 3},
		},
		{
			name: "lc case 3",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOrderPerOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: []int{0, 1},
		},
		{
			name: "lc case 2",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0}, {2, 0}, {3, 1}, {3, 2},
				},
			},
			want: []int{0, 2, 1, 3},
		},
		{
			name: "lc case 3",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrderPerOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrderPerOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOrderDFS(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: []int{0, 1},
		},
		{
			name: "lc case 2",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0}, {2, 0}, {3, 1}, {3, 2},
				},
			},
			want: []int{0, 2, 1, 3},
		},
		{
			name: "lc case 3",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrderDFS(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrderDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOrderBFSPV1(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "lc case 1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: []int{0, 1},
		},
		{
			name: "lc case 2",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0}, {2, 0}, {3, 1}, {3, 2},
				},
			},
			want: []int{0, 2, 1, 3},
		},
		{
			name: "lc case 3",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrderBFSPV1(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrderBFSPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
