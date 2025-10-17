package dp

func minInsertions(s string) int {
	return minInsertionsDPT(s)
}

func minInsertionsDPT(s string) int {
	return len(s) - longestPalindromeSubseqDPT(s)
}
