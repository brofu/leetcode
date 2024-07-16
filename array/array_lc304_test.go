package array

import (
	"fmt"
	"testing"
)

func TestNumMatrix_SumRegion(t *testing.T) {
	type fields struct {
		matrix [][]int
	}
	type args struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "lc case 1",
			fields: fields{
				matrix: [][]int{
					{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5},
				},
			},
			args: args{
				row1: 2,
				col1: 1,
				row2: 4,
				col2: 3,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := ConstructorNumMatrix(tt.fields.matrix)
			fmt.Println(this.preSum)
			if got := this.SumRegion(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2); got != tt.want {
				t.Errorf("NumMatrix.SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
