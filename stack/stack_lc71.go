package stack

import (
	"bytes"
	"strings"
)

func simplifyPath(path string) string {

	buf := bytes.Buffer{}
	st := make([]string, 0)

	pathes := strings.Split(path, "/")

	for _, p := range pathes {
		if p == "" || p == "." {
			continue
		}
		if p == ".." {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
			continue
		}
		st = append(st, p)
	}
	for _, p := range st {
		buf.WriteString("/")
		buf.WriteString(p)
	}

	r := buf.String()
	if r == "" {
		r = "/"
	}
	return r
}
