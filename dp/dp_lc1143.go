package dp

import "github.com/brofu/leetcode/common"

/**
DP function version. Top to Down approach

KP.	How to define the DP function?
*/
func longestCommonSubsequence(text1 string, text2 string) int {

	// setup the memo
	memo := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		memo[i] = make([]int, len(text2))
		for j := 0; j < len(text2); j++ {
			memo[i][j] = -1
		}
	}

	var dp func(string, int, string, int) int

	// dp defination: the LCS of s1[i...] and s2[j...]
	dp = func(s1 string, i int, s2 string, j int) int {

		// base case
		if i == len(s1) || j == len(s2) {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if s1[i] == s2[j] {
			memo[i][j] = 1 + dp(s1, i+1, s2, j+1)
		} else {
			memo[i][j] = common.MaxInt(
				dp(s1, i+1, s2, j), //s1[i] not in lcs
				dp(s1, i, s2, j+1), // s2[j] not in lcs
			)
		}
		return memo[i][j]
	}

	return dp(text1, 0, text2, 0)
}

/**
DP table version. Down to Top approach
KP.	How to define the DP table?
*/
func longestCommonSubsequenceDPTable(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// dp defination: dp[i][j] is the length of LCS of text1[0...i] and text2[0...j]
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// base case
	// dp[0][0...n] = dp[0...m][0] = 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = common.MaxInt(
					dp[i-1][j], // text1[i] not in lcs
					dp[i][j-1], // text2[j] not in lcs
				)
			}
		}
	}

	return dp[m][n]
}

/**
DP table version with space compression.
KP.
	1.	Similar logic as DP table, but reuse the 1-D DP table.
	2.	Need to consider how to store the 3 values:
		a.	dp[i-1][j-1], // use a `tp` var
		b.	dp[i][j-1],   // for a same is, dp[i][j-1] actually is dp[j-1]
		c.	dp[i-1][j],   // for a `i`, BEFORE dp[j] is updated, it's actually dp[i-1][p]
*/

func longestCommonSubsequenceDPTableWithSpaceCompress(text1 string, text2 string) int {

	dp := make([]int, len(text2)+1)

	for i := 1; i <= len(text1); i++ {
		tp := dp[0]

		for j := 1; j <= len(text2); j++ {

			if text1[i-1] == text2[j-1] { // equal
				tmp := dp[j]   // dp[j] = dp[i-1][j] before update
				dp[j] = tp + 1 // tp is dp[i-1][j-1]
				tp = tmp
				continue
			}
			// not equal
			tmp := dp[j]
			dp[j] = common.MaxInt(
				dp[j],   // dp[j] acutally is dp[i-1][j]
				dp[j-1], // dp[j-1] actually is dp[i][j-1]
			)
			tp = tmp

		}
	}

	return dp[len(text2)]
}
