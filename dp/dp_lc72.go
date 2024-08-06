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

func minDistanceDPTableWithSpaceCompress(word1 string, word2 string) int {

	//TODO
	return 0
}
