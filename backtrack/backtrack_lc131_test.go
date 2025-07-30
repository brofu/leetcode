package backtrack

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case 1",
			args: args{
				s: "aab",
			},
			want: [][]string{
				{"a", "a", "b"}, {"aa", "b"},
			},
		},
		{
			name: "case 2",
			args: args{
				s: "efe",
			},
			want: [][]string{
				{"e", "f", "e"}, {"efe"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partitionWithCache(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case 1",
			args: args{
				s: "aab",
			},
			want: [][]string{
				{"a", "a", "b"}, {"aa", "b"},
			},
		},
		{
			name: "case 2",
			args: args{
				s: "efe",
			},
			want: [][]string{
				{"e", "f", "e"}, {"efe"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionWithCache(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionWithCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partitionWithCacheV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case 1",
			args: args{
				s: "aab",
			},
			want: [][]string{
				{"a", "a", "b"}, {"aa", "b"},
			},
		},
		{
			name: "case 2",
			args: args{
				s: "efe",
			},
			want: [][]string{
				{"e", "f", "e"}, {"efe"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionWithCacheV2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionWithCacheV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
