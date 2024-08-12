package dp

import "github.com/brofu/leetcode/common"

/**
KP.
	1.	The result is actually the `LCS` of 2 strings
	2.	So the problem is actually to calculate the `LCS` of these 2 strings
*/
func minDistance583(word1 string, word2 string) int {

	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = common.MaxInt(
					dp[i-1][j],
					dp[i][j-1],
				)
			}
		}
	}

	return m + n - 2*dp[m][n]
}

func minDistance583WithSpaceCompress(word1 string, word2 string) int {

	dp := make([]int, len(word2)+1)

	for i := 1; i <= len(word1); i++ {
		ph := dp[0]

		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				tmp := dp[j]
				dp[j] = 1 + ph // ph is actually dp[i-1][j-1], why?
				ph = tmp       // record new dp[i-1][j-1]
			} else {
				tmp := dp[j]
				dp[j] = common.MaxInt(
					dp[j],   // dp[j] is actually dp[i-1][j], before it's updated
					dp[j-1], // dp[j-1] is actually dp[i][j-1]
				)
				ph = tmp // now, for new j with 1 increased, ph is actually dp[i-1][j-1]
			}
		}
	}

	return len(word1) + len(word2) - 2*dp[len(word2)]
}
