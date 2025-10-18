package dp

import (
	"github.com/brofu/leetcode/common"
)

func minDistance(word1 string, word2 string) int {
	//return minDistanceDPF20251018(word1, word2)
	return minDistanceDPT20251018(word1, word2)
}

/*
*
DP function version. Top to Down approach
*/
func minDistanceV0(word1 string, word2 string) int {
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

/*
*
DP table version with space compress.
It's tricky to understand.

KP.
 1. The overall logic is same as the DP table version, but need to reuse dp table for each `i` for `word1`
 2. In the inner loop at each `j` on `word2`,
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

func minDistanceDPTablePV1(word1 string, word2 string) int {

	m, n := len(word1), len(word2)

	// setup 2-D DP table
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + common.MinIntMultiple(
					dp[i-1][j],   // delete
					dp[i][j-1],   // insert
					dp[i-1][j-1], // replacement
				)
			}
		}
	}

	return dp[m][n]
}

func minDistanceDPFPV2(word1 string, word2 string) int {

	m, n := len(word1)-1, len(word2)-1
	memo := make([][]int, m+1)
	for idx := range memo {
		memo[idx] = make([]int, n+1)
		for iidx := range memo[idx] {
			memo[idx][iidx] = -1
		}
	}

	var dp func(int, int) int
	dp = func(i, j int) int {

		//base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// recursive
		var result int
		if word1[i] == word2[j] {
			result = dp(i-1, j-1)
		} else {
			result = common.MinIntThree(dp(i-1, j), dp(i-1, j-1), dp(i, j-1)) + 1
		}
		memo[i][j] = result

		return result
	}

	result := dp(m, n)
	return result
}

func minDistanceDPTablePV2(word1 string, word2 string) int {

	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}
	for idx := range dp[0] {
		dp[0][idx] = idx
	}
	for idx := 0; idx <= m; idx++ {
		dp[idx][0] = idx
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.MinIntThree(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

/*
DP function

TC:
1. setup memeo O(m*n)
2. Recursive depth, roughly O(m*n), with memo

SC:
1. memo O(m*n)
2. Recursive stack O(m*n)
*/
func minDistanceDPF20251018(word1 string, word2 string) int {

	m, n := len(word1), len(word2)

	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// the min distance between word1[0...i] and world[0...j]
	var dp func(i, j int) int

	dp = func(i, j int) int {
		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// state transfer
		if word1[i] == word2[j] {
			memo[i][j] = dp(i-1, j-1)
		} else {
			memo[i][j] = common.MinIntThree(dp(i, j-1)+1, dp(i-1, j)+1, dp(i-1, j-1)+1)
		}
		return memo[i][j]
	}

	return dp(m-1, n-1)
}

/*
DP Table

TC:
1. setup the DP table O(m*n)
2. Fill up the table O(m*n)

SC:
1. DP table O(m*n). Can compress.
*/

func minDistanceDPT20251018(word1 string, word2 string) int {

	m, n := len(word1), len(word2)

	// dp[i][j] means the min distance between word1[0...i-1], word2[0...j-1]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.MinIntThree(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[m][n]
}
