package dp

import (
	"github.com/brofu/leetcode/common"
)

/**
DP function version. Top to Down approach
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1)-1, len(word2)-1

	// set up the memo
	memo := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}

	var dp func(string, int, string, int) int

	dp = func(s1 string, i int, s2 string, j int) int {

		// base case
		if i == -1 {
			return j + 1
		}

		// base case
		if j == -1 {
			return i + 1
		}

		// existing in memo
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if s1[i] == s2[j] { // equal, do nothing
			memo[i][j] = dp(s1, i-1, s2, j-1)
			return memo[i][j]
		} else {
			memo[i][j] = common.MinIntMultiple(
				dp(s1, i, s2, j-1)+1,   // insert to s1
				dp(s1, i-1, s2, j)+1,   // delete in s1
				dp(s1, i-1, s2, j-1)+1, // replace s1
			)
			return memo[i][j]
		}

	}

	return dp(word1, m, word2, n)
}

func minDistanceDPTable(word1 string, word2 string) int {

	m, n := len(word1), len(word2)

	// setup up DP table
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.MinIntMultiple(
					dp[i-1][j]+1,
					dp[i][j-1]+1,
					dp[i-1][j-1]+1,
				)
			}
		}
	}

	return dp[m][n]
}

/**
DP table version with space compress.
It's tricky to understand.

KP.
	1.	The overall logic is same as the DP table version, but need to reuse dp table for each `i` for `word1`
	2.	In the inner loop at each `j` on `word2`,
		a.	before dp[j] is updated, dp[j] is actually dp[i-1][j], and
		b.	after dp[j] is updated, dp[j] is actually dp[i][j]
*/
func minDistanceDPTableWithSpaceCompress(word1 string, word2 string) int {

	dp := make([]int, len(word2)+1)

	for i := range dp {
		dp[i] = i
	}

	for i := 1; i <= len(word1); i++ {
		topLeft := dp[0] // KP1: topLeft actually is dp[i-1][0]
		dp[0]++

		for j := 1; j <= len(word2); j++ {

			// KP2:
			// this is tricky
			// if i == 1, actually, topLeft is dp[i-1][0]
			// if i != 1, actually, topLeft is dp[i-1][j-1], why? Refer to KP3, KP4, and KP5
			if word1[i-1] == word2[j-1] {
				dp[j], topLeft = topLeft, dp[j]
				// same code, better to understand
				// tmp := dp[j]
				// dp[j] = topLeft
				// topLeft = tmp
				continue
			}

			tmp := dp[j] // KP3: before dp[j] is updated, it's actually dp[i-1][j]
			dp[j] = 1 + common.MinIntMultiple(
				dp[j-1], // insert
				dp[j],   //delete
				topLeft, //replacement
			) // KP4: after dp[j] is updated, it's dp[i][j]
			topLeft = tmp // KP5: topLeft here actually is dp[i-1][j]
		}
	}
	return dp[len(word2)]
}
