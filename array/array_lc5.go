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
