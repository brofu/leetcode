package dp

import (
	"github.com/brofu/leetcode/common"
)

func longestPalindromeSubseq(s string) int {

	//return longestPalindromeSubseqDPF(s)
	return longestPalindromeSubseqDPT(s)

}

func longestPalindromeSubseqDPF(s string) int {

	n := len(s)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// dp(i, j): the longest palindrome sub squence of d[i...j]
	var dp func(int, int) int

	dp = func(i, j int) int {

		// base case
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if s[i] == s[j] {
			memo[i][j] = dp(i+1, j-1) + 2
		} else {
			memo[i][j] = common.MaxInt(dp(i, j-1), dp(i+1, j))
		}

		return memo[i][j]
	}

	return dp(0, n-1)
}

func longestPalindromeSubseqDPT(s string) int {

	n := len(s)
	// dp[i][j]: longest palindrome sub sequence between s[i...j]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // if i == j, dp[i][j] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = common.MaxInt(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]
}
