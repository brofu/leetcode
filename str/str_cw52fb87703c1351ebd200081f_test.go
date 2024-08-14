package str

import "testing"

func TestWhatCentury(t *testing.T) {
	type args struct {
		year string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "cw case 1",
			args: args{
				year: "1999",
			},
			want: "20th",
		},
		{
			name: "cw case 2",
			args: args{
				year: "2011",
			},
			want: "21st",
		},
		{
			name: "cw case 3",
			args: args{
				year: "2154",
			},
			want: "22nd",
		},
		{
			name: "cw case 4",
			args: args{
				year: "2259",
			},
			want: "23rd",
		},
		{
			name: "cw case 5",
			args: args{
				year: "1124",
			},
			want: "12th",
		},
		{
			name: "cw case 6",
			args: args{
				year: "2000",
			},
			want: "20th",
		},
		{
			name: "cw case 7",
			args: args{
				year: "0000",
			},
			want: "0th",
		},
		{
			name: "cw case 8",
			args: args{
				year: "9999",
			},
			want: "100th",
		},
		{
			name: "cw case 9",
			args: args{
				year: "0900",
			},
			want: "9th",
		},
		{
			name: "cw case 9",
			args: args{
				year: "0001",
			},
			want: "1st",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhatCentury(tt.args.year); got != tt.want {
				t.Errorf("WhatCentury() = %v, want %v", got, tt.want)
			}
		})
	}
}
