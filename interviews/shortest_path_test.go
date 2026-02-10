package interviews

import (
	"reflect"
	"testing"
)

func Test_findPath(t *testing.T) {
	type args struct {
		input       []string
		source      string
		destination string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				input: []string{
					"NY : SG",
					"NY : LD",
					"LD : NY",
					"LD : BK",
					"SG : TK",
					"TK : BK",
				},
				source:      "NY",
				destination: "BK",
			},
			want: "NY : LD : BK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("input", tt.args.input)
			if got := findPath(tt.args.input, tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("findPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPathVar(t *testing.T) {
	type args struct {
		input       []string
		source      string
		destination string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				input: []string{
					"NY : SG",
					"NY : LD",
					"NY : LD2",
					"LD : NY",
					"LD : BK",
					"LD2 : BK",
					"SG : TK",
					"TK : BK",
				},
				source:      "NY",
				destination: "BK",
			},
			want: []string{"NY : LD : BK", "NY : LD2 : BK"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("input", tt.args.input)
			if got := findPathVar(tt.args.input, tt.args.source, tt.args.destination); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPathVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPathVarDFS(t *testing.T) {
	type args struct {
		input       []string
		source      string
		destination string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				input: []string{
					"NY : SG",
					"NY : LD",
					"NY : LD2",
					"LD : NY",
					"LD : BK",
					"LD2 : BK",
					"SG : TK",
					"TK : BK",
				},
				source:      "NY",
				destination: "BK",
			},
			want: []string{"NY : LD : BK", "NY : LD2 : BK"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPathVarDFS(tt.args.input, tt.args.source, tt.args.destination); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPathVarDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
