package dp

import (
	"github.com/brofu/leetcode/common"
)

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
