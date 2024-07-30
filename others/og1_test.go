package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_resolve(t *testing.T) {
	type args struct {
		input []Record
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal case 1",
			args: args{
				input: []Record{
					{"PZ", 30}, {"PZ", 47},
				},
			},
			want: true,
		},
		{
			name: "normal case 2",
			args: args{
				input: []Record{
					{"MD", 30}, {"CV", 40}, {"MD", 37},
				},
			},
			want: true,
		},
		{
			name: "normal case 3",
			args: args{
				input: []Record{
					{"JJ", 1},
				},
			},
			want: true,
		},
		{
			name: "normal case 4",
			args: args{
				input: []Record{
					{"MD", 1},
				},
			},
			want: false,
		},
		{
			name: "normal case 5",
			args: args{
				input: []Record{},
			},
			want: false,
		},
		{
			name: "normal case 6",
			args: args{
				input: []Record{
					{"PZ", 30}, {"PZ", 73}, {"MD", 48}, {"MD", 105},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolve(tt.args.input); got != tt.want {
				t.Errorf("resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolveV2(t *testing.T) {
	type args struct {
		input []Record
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{
			name: "normal case 1",
			args: args{
				input: []Record{
					{"PZ", 15}, {"PZ", 32},
				},
			},
			want:  "vaccinated",
			want1: 46,
		},
		{
			name: "normal case 2",
			args: args{
				input: []Record{
					{"MD", 0}, {"MD", 24}, {"PZ", 25}, {"PZ", 42},
				},
			},
			want:  "vaccinated",
			want1: 38,
		},
		{
			name: "normal case 3",
			args: args{
				input: []Record{
					{"MD", 30}, {"MD", 40}, {"MD", 54},
				},
			},
			want:  "not_vaccinated",
			want1: -1,
		},
		{
			name: "normal case 4",
			args: args{
				input: []Record{
					{"JJ", 10}, {"CV", 11}, {"CV", 24},
				},
			},
			want:  "vaccinated",
			want1: 24,
		},
		{
			name: "normal case 5",
			args: args{
				input: []Record{
					{"CV", 10}, {"CV", 35}, {"MD", 20},
				},
			},
			want:  "not_vaccinated",
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := resolveV2(tt.args.input)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, got1, tt.want1)
		})
	}
}
