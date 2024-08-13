package dp

import (
	"github.com/brofu/leetcode/common"
)

/**
KP.
	1.	Similar as 583 and 1143. But cannot get resolved by reusing LCS
	2.	Return the min sum from DP function.
*/
func minimumDeleteSum(s1 string, s2 string) int {

	memo := make([][]int, len(s1))
	for i := range memo {
		memo[i] = make([]int, len(s2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(int, int) int

	dp = func(i int, j int) int {

		// base case 1
		if i == len(s1) {
			result := 0
			for ; j < len(s2); j++ {
				result += int(s2[j])
			}
			return result
		}
		// base case 2
		if j == len(s2) {
			result := 0
			for ; i < len(s1); i++ {
				result += int(s1[i])
			}
			return result
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if s1[i] == s2[j] {
			memo[i][j] = dp(i+1, j+1)
		} else {
			memo[i][j] = common.MinInt(
				int(s1[i])+dp(i+1, j), // delete s1[i]
				int(s2[j])+dp(i, j+1), // delete s2[j]
			)
		}
		return memo[i][j]
	}

	return dp(0, 0)
}

/**
KP.
	1.	Similar as 583 and 1143. But cannot get resolved by reusing LCS
	2.	Record the min sum in DP table
*/

func minimumDeleteSumDPTable(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// base case 1
	for i := 1; i <= m; i++ {
		dp[i][0] = int(s1[i-1]) + dp[i-1][0]
	}

	// base case 2
	for j := 1; j <= n; j++ {
		dp[0][j] = int(s2[j-1]) + dp[0][j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.MinInt(
					int(s1[i-1])+dp[i-1][j], // delete s1[i]
					int(s2[j-1])+dp[i][j-1], // delete s2[j]
				)
			}
		}
	}

	return dp[m][n]
}

/**
DP table version, with space compress
*/
func minimumDeleteSumDPTableWithSpaceCompress(s1 string, s2 string) int {

	dp := make([]int, len(s2)+1)

	// base case
	for j := 1; j <= len(s2); j++ {
		dp[j] = dp[j-1] + int(s2[j-1])
	}

	for i := 1; i <= len(s1); i++ {
		ph := dp[0]
		dp[0] += int(s1[i-1])

		for j := 1; j <= len(s2); j++ {
			tmp := dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = ph
			} else {
				dp[j] = common.MinInt(
					int(s1[i-1])+dp[j],   // delete s1[i]. dp[j] actually dp[i-1][j] before updated
					int(s2[j-1])+dp[j-1], // delete s2[j], dp[j-1] actually dp[i][j-1]
				)
			}
			ph = tmp // ph is actually dp[i-1][j-1], for the next new j
		}

	}
	return dp[len(s2)]
}
