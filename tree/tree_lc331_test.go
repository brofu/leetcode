package tree

import (
	"testing"
)

func Test_isValidSerialization(t *testing.T) {
	type args struct {
		preorder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				preorder: "1,#",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				preorder: "9,#,#,1",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				preorder: "1",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				preorder: "9,#,92,#,#",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSerialization(tt.args.preorder); got != tt.want {
				t.Errorf("isValidSerialization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSerializationSubTask(t *testing.T) {
	type args struct {
		preorder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				preorder: "1,#",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				preorder: "9,#,#,1",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				preorder: "1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSerializationSubTask(tt.args.preorder); got != tt.want {
				t.Errorf("isValidSerializationSubTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
