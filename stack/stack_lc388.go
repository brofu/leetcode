package stack

import (
	"strings"

	"github.com/brofu/leetcode/common"
)

/*
kP:
	1. Use `\t` as layer flag, and
	2. Control the stack with it.
	3. Usage of `strings.LastIndexByte`
*/
func lengthLongestPath(input string) int {

	result := 0
	st := make([]int, 0)

	pathes := strings.Split(input, "\n")

	for _, path := range pathes {

		level := strings.LastIndexByte(path, '\t') + 1

		for len(st) > level {
			st = st[:len(st)-1]
		}

		st = append(st, len(path[level:])) // length of current path

		if strings.Contains(path, ".") { //file
			t := 0 // the length of names
			for _, c := range st {
				t += c
			}
			result = common.MaxInt(result, t+len(st)-1) // length of names + number of "/"
		}
	}

	return result
}
