package fs

import (
	"fmt"
	"testing"
)

func TestFileSystemInMemory(t *testing.T) {
	tests := []struct {
		name    string
		f       func()
		enabled bool
	}{
		{
			name: "lc case 1",
			f: func() {
				fs := Constructor()
				fmt.Println(fs.Ls("/"))
				fs.Mkdir("/a/b/c")
				fmt.Println(fs.Ls("/"))
				fs.AddContentToFile("/a/b/c/d", "hello")
				fmt.Println(fs.Ls("/"))
				fmt.Println(fs.ReadContentFromFile("/a/b/c/d"))
			},
		},
		{
			name:    "lc case 2",
			enabled: false,
			f: func() {
				fs := Constructor()
				fmt.Println(fs.Ls("/"))
				fs.Mkdir("/a/b/c")
				fmt.Println(fs.Ls("/a/b"))
			},
		},
		{
			name:    "lc case 3",
			enabled: false,
			f: func() {
				fs := Constructor()
				fmt.Println(fs.Ls("/"))
				fs.Mkdir("/a/b/c")
				fmt.Println(fs.Ls("/a/b"))
				fs.Mkdir("/a/b/a")
				fmt.Println(fs.Ls("/a/b"))
			},
		},
		{
			name:    "lc case 4",
			enabled: false,
			f: func() {
				fs := Constructor()
				fmt.Println(fs.Ls("/"))
				fs.Mkdir("/goowmfn")
				fmt.Println(fs.Ls("/"))
				fs.Mkdir("/z")
				fmt.Println(fs.Ls("/"))
				fs.AddContentToFile("/goowmfn/c", "shetopcy")
			},
		},
		{
			name:    "lc case 5",
			enabled: false,
			f: func() {
				fs := Constructor()
				fmt.Println(fs.Ls("/"))
				fs.Mkdir("/lku")
				fmt.Println(fs.Ls("/lku"))
				fs.Mkdir("/jgwwxo")
				fmt.Println(fs.Ls("/"))
				fs.AddContentToFile("/rce", "me")
				fmt.Println(fs.Ls("/"))
			},
		},
		{
			name:    "lc case 6",
			enabled: true,
			f: func() {
				fs := Constructor()
				fmt.Println(fs.Ls("/"))
				fs.Mkdir("/i")
				fmt.Println(fs.Ls("/lku"))
				fs.Mkdir("/uqfn")
				fmt.Println(fs.Ls("/"))
				fs.AddContentToFile("/ooosoc", "cv")
				fmt.Println(fs.Ls("/"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.enabled {
				tt.f()
			}
		})
	}

}
