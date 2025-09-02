package dp

import (
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("name: %s, lengthOfLIS() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISWithBS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISWithBS(tt.args.nums); got != tt.want {
				t.Errorf("name: %s, lengthOfLISWithBS() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISWithBSPV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0},
			},
			want: 1,
		},
		{
			name: "lc case 5",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISWithBSPV1(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLISWithBSPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISPV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISPV1(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLISPV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISPV2DPF(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.nums)
			if got := lengthOfLISPV2DPF(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLISPV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISPV2DPA(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		},
		{
			name: "lc case 5",
			args: args{
				nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISPV2DPA(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLISPV2DPA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISWithBSPV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		},
		{
			name: "lc case 5",
			args: args{
				nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISWithBSPV2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLISWithBSPV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISWithBSPV3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "lc case 4",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		},
		{
			name: "lc case 5",
			args: args{
				nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.nums)
			if got := lengthOfLISWithBSPV3(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLISWithBSPV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
