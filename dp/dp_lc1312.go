package dp

/*
Very similar as problem 516,
After getting the longest palindrome sub sequence, can get the numbers need to insert or delete
*/

func minInsertions(s string) int {
	return minInsertionsDPT(s)
}

func minInsertionsDPT(s string) int {
	return len(s) - longestPalindromeSubseqDPT(s)
}
