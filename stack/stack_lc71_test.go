package stack

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "case 2",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			name: "case 3",
			args: args{
				path: "/home/user/Documents/../Pictures",
			},
			want: "/home/user/Pictures",
		},
		{
			name: "case 4",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "case 5",
			args: args{
				path: "/.../a/../b/c/../d/./",
			},
			want: "/.../b/d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
