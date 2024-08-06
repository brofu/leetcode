package others

import "github.com/brofu/leetcode/tree/trietree"

func resovle(s []string) string {

	if len(s) == 0 {
		return ""
	}

	result := ""

	for index, r := range s[0] {

		for i := 1; i < len(s); i++ {

			ss := s[i]

			if len(ss) < index+1 {
				return result
			}

			if ss[index] != byte(r) {
				return result
			}
		}

		result += string(r)
	}

	return result
}

func resovleV2(s []string) string {
	if len(s) == 0 {
		return ""
	}
	tm := trietree.NewTrieMap()
	for _, w := range s {
		tm.Put(w, struct{}{})
	}
	return tm.LongestPrefixV2()
}
