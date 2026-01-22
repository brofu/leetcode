package dp

import (
	"testing"
)

func Test_maxEnvelopesV2(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				envelopes: [][]int{
					{5, 4}, {6, 4}, {6, 7}, {2, 3},
				},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				envelopes: [][]int{
					{1, 1}, {1, 1}, {1, 1},
				},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				envelopes: [][]int{
					{30, 50}, {12, 2}, {3, 4}, {12, 15},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.envelopes)
			if got := maxEnvelopesV2(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxEnvelopesDP(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				envelopes: [][]int{
					{5, 4}, {6, 4}, {6, 7}, {2, 3},
				},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				envelopes: [][]int{
					{1, 1}, {1, 1}, {1, 1},
				},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				envelopes: [][]int{
					{30, 50}, {12, 2}, {3, 4}, {12, 15},
				},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				envelopes: [][]int{
					{5, 4}, {5, 2},
				},
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				envelopes: [][]int{
					{5, 2}, {5, 4},
				},
			},
			want: 1,
		},
		{
			name: "case 6",
			args: args{
				envelopes: [][]int{
					{2, 3},
					{5, 4},
					{6, 7},
					{6, 4}, // same width as {6, 7}, smaller height
					{7, 5},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopesDP(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopesDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxEnvelopesV3(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				envelopes: [][]int{
					{5, 4}, {6, 4}, {6, 7}, {2, 3},
				},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				envelopes: [][]int{
					{1, 1}, {1, 1}, {1, 1},
				},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				envelopes: [][]int{
					{30, 50}, {12, 2}, {3, 4}, {12, 15},
				},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				envelopes: [][]int{
					{5, 4}, {5, 2},
				},
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				envelopes: [][]int{
					{5, 2}, {5, 4},
				},
			},
			want: 1,
		},
		{
			name: "case 6",
			args: args{
				envelopes: [][]int{
					{2, 3},
					{5, 4},
					{6, 7},
					{6, 4}, // same width as {6, 7}, smaller height
					{7, 5},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopesV3(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopesV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxEnvelopesV4(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name    string
		enabled bool
		args    args
		want    int
	}{
		{
			name: "case 1",
			args: args{
				envelopes: [][]int{
					{5, 4}, {6, 4}, {6, 7}, {2, 3},
				},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				envelopes: [][]int{
					{1, 1}, {1, 1}, {1, 1},
				},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				envelopes: [][]int{
					{30, 50}, {12, 2}, {3, 4}, {12, 15},
				},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				envelopes: [][]int{
					{5, 4}, {5, 2},
				},
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				envelopes: [][]int{
					{5, 2}, {5, 4},
				},
			},
			want: 1,
		},
		{
			name: "case 6",
			args: args{
				envelopes: [][]int{
					{2, 3},
					{5, 4},
					{6, 7},
					{6, 4}, // same width as {6, 7}, smaller height
					{7, 5},
				},
			},
			want: 3,
		},
		{
			name:    "case 7",
			enabled: true,
			args: args{
				envelopes: [][]int{
					{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		if !tt.enabled {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopesV4(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopesV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
