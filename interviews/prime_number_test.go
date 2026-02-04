package interviews

import (
	"testing"
)

func TestCheckPrimeNumberIndex(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				number: 11,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPrimeNumberIndex(tt.args.number); got != tt.want {
				t.Errorf("CheckPrimeNumberIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPrimeNumberIndexV2(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				number: 11,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPrimeNumberIndexV2(tt.args.number); got != tt.want {
				t.Errorf("CheckPrimeNumberIndexV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPrimeNumberIndexV3(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				number: 11,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPrimeNumberIndexV3(tt.args.number); got != tt.want {
				t.Errorf("CheckPrimeNumberIndexV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
