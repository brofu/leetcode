package array

/**
KP.	找回文串的难点在于，回文串的的长度可能是奇数也可能是偶数，核心是从中心向两端扩散的双指针技巧。
*/
func longestPalindrome(s string) string {

	result := ""

	for i := 0; i < len(s); i++ {
		s1 := findPalindrome(s, i, i)
		s2 := findPalindrome(s, i, i+1)
		if len(result) < len(s1) {
			result = s1
		}
		if len(result) < len(s2) {
			result = s2
		}
	}
	return result
}

func findPalindrome(s string, left, right int) string {

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	// note: s[len(s):len(s)] would not report error
	return s[left+1 : right]
}

/*
KP:
	1. Why this is wrong? Only consider ODD scenario
*/
func longestPalindromePV2(s string) string {

	result := ""
	for idx := 0; idx < len(s); idx++ {
		r1 := findTheString(s, idx, idx)
		r2 := findTheString(s, idx, idx+1)
		if len(result) < len(r1) {
			result = r1
		}
		if len(result) < len(r2) {
			result = r2
		}
	}

	return result
}

func findTheString(s string, start, end int) string {
	for ; start >= 0 && end < len(s); start, end = start-1, end+1 {
		if s[start] != s[end] {
			break
		}
	}

	return s[start+1 : end]
}
