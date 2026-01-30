package dp

import (
	"testing"
)

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 5, 11, 5},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartitionWithSpaceCompression(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 5, 11, 5},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1, 2, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionWithSpaceCompression(tt.args.nums); got != tt.want {
				t.Errorf("canPartitionWithSpaceCompression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartitionV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 5, 11, 5},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1, 2, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionDPV1(tt.args.nums); got != tt.want {
				t.Errorf("canPartitionV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartitionDPCompressedV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				nums: []int{1, 5, 11, 5},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
		{
			name: "lc case 3",
			args: args{
				nums: []int{1, 2, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionDPCompressedV1(tt.args.nums); got != tt.want {
				t.Errorf("canPartitionDPCompressedV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
