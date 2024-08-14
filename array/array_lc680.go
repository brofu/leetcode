package array

/**
Wrong Version
KP.
	1.	Cannot check only by current indexed character s[i] and s[j]
	2.	Instead, need to check all the left string
*/
func validPalindromeWrongVersion(s string) bool {

	count := 0

	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			hit := false
			if s[i] == s[j-1] {
				j--
				hit = true
			} else if s[i+1] == s[j] {
				i++
				hit = true
			}

			if !hit {
				return false
			}

			count++
			if count > 1 {
				break
			}
		}
	}

	return count <= 1
}

func validPalindrome(s string) bool {

	helper := func(i, j int) bool {
		for ; i <= j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return helper(i+1, j) || helper(i, j-1) // check if s1[i+1...j] or s1[i...j-1] is palindrome
		}
	}

	return true
}
