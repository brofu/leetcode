package str

import (
	"bytes"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	minLength := 200

	for _, str := range strs {
		if minLength > len(str) {
			minLength = len(str)
		}
	}

	if minLength == 0 {
		return ""
	}

	buf := bytes.Buffer{}

stop:
	for i := 0; i < minLength; i++ {
		p := strs[0][i]
		j := 1
		for ; j < len(strs); j++ {
			if p != strs[j][i] {
				break stop
			}
		}
		if j == len(strs) {
			buf.WriteByte(p)
		}
	}

	return buf.String()
}

/*

A more brief version
*/
func longestCommonPrefixV2(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	minLength := 200

	for _, str := range strs {
		if minLength > len(str) {
			minLength = len(str)
		}
	}

	if minLength == 0 {
		return ""
	}

	for i := 0; i < minLength; i++ {
		p := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if p != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:minLength]
}
