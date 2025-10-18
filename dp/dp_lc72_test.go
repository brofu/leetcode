package dp

import (
	"testing"
)

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistanceV0(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceV0(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistanceDPTable(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lc case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceDPTable(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistanceDPTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistanceDPTableWithSpaceCompress(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "self case 1",
			args: args{
				word1: "rat",
				word2: "ras",
			},
			want: 1,
		},

		{
			name: "lc case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceDPTableWithSpaceCompress(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistanceDPTableWithSpaceCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistanceDPTablePV1(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "self case 1",
			args: args{
				word1: "rat",
				word2: "ras",
			},
			want: 1,
		},
		{
			name: "lc case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceDPTablePV1(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistanceDPTablePV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistanceDPFPV2(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name    string
		enabled bool
		args    args
		want    int
	}{
		{
			name:    "self case 1",
			enabled: true,
			args: args{
				word1: "rat",
				word2: "ras",
			},
			want: 1,
		},
		{
			name:    "lc case 2",
			enabled: true,
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name:    "lc case 3",
			enabled: true,
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		if !tt.enabled {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceDPFPV2(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistanceDPFPV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistanceDPTablePV2(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "self case 1",
			args: args{
				word1: "rat",
				word2: "ras",
			},
			want: 1,
		},
		{
			name: "lc case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "lc case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceDPTablePV2(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistanceDPTablePV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistanceDPF20251018(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistanceDPF20251018(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistanceDPF20251018() = %v, want %v", got, tt.want)
			}
		})
	}
}
