package stream

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want float64
	}{
		{
			name: "lc case 1",
			nums: []int{1, 2},
			want: 1.5,
		},
		{
			name: "lc case 2",
			nums: []int{1, 2, 3},
			want: 2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor()
			for _, num := range tt.nums {
				obj.AddNum(num)
			}
			got := obj.FindMedian()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Result = %v, want %v", got, tt.want)
			}
		})
	}
}
