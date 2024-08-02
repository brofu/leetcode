package array

/**
KP.
	1.	Get the idea from problem 5. Longest Palindromic Substring
	2.	If we get the longest palindromic which MUST start with the 1st character of s, then,
	3.	We can just add the string of `Reverse(s[longest:len(s)])` to the original string `s`
*/
func shortestPalindrome(s string) string {
	loggest := 0
	differS := ""

	for i := 0; i < len(s); i++ {
		l1, r1 := findPalindromeV2(s, i, i)
		l2, r2 := findPalindromeV2(s, i, i+1)

		if loggest < r1-l1 && l1 == 0 {
			loggest = r1 - l1
		}
		if loggest < r2-l2 && l2 == 0 {
			loggest = r2 - l2
		}
	}

	for i := len(s) - 1; i >= loggest; i-- {
		differS += string(s[i])
	}

	return differS + s
}

func findPalindromeV2(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right
}
