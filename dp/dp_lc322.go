package dp

import (
	"math"

	"github.com/brofu/leetcode/common"
)

/**
KP
	1.	最优子结构
	2.	状态转移方程
	3.  重叠子问题
*/
func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	min := math.MaxInt

	for _, coin := range coins {

		subproblem := coinChange(coins, amount-coin)

		if subproblem == -1 {
			continue
		}

		min = MinInt(subproblem+1, min)
	}

	if min == math.MaxInt {
		min = -1
	}

	return min
}

func coinChangeRecursiveWithMemo(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for index, _ := range memo {
		memo[index] = 0
	}

	return coinChangeWithMemoWrapper(memo, coins, amount)
}

func coinChangeWithMemoWrapper(memo []int, coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if memo[amount] != 0 {
		return memo[amount]
	}

	max := amount + 1
	min := max

	for _, coin := range coins {
		subproblem := coinChangeWithMemoWrapper(memo, coins, amount-coin)

		if subproblem == -1 {
			continue
		}

		min = MinInt(subproblem+1, min)
	}

	if min == max {
		min = -1
	}

	memo[amount] = min

	return min
}

func coinChangeRecurrenceWithMemo(coins []int, amount int) int {

	extraLength := amount + 1

	memo := make([]int, extraLength)
	for index, _ := range memo {
		memo[index] = extraLength
	}
	memo[0] = 0

	for i := 1; i < extraLength; i++ {
		for _, coin := range coins {
			if i-coin >= 0 { // no result
				memo[i] = MinInt(memo[i], memo[i-coin]+1)
			}
		}
	}

	if memo[amount] == extraLength {
		memo[amount] = -1
	}

	return memo[amount]
}

func coinChangePV1(coins []int, amount int) int {

	var dp func(int) int

	memo := make(map[int]int)

	dp = func(amount int) int {

		// base case
		if amount == 0 {
			return 0
		}

		// wrong solution
		if amount < 0 {
			return -1
		}

		if val, exists := memo[amount]; exists {
			return val
		}

		solution := math.MaxInt
		for _, coin := range coins {
			subSolution := dp(amount - coin)
			if subSolution == -1 {
				continue
			}
			solution = MinInt(solution, subSolution+1)
		}

		if solution == math.MaxInt {
			solution = -1
		}

		memo[amount] = solution
		return solution
	}

	return dp(amount)
}

/**
KP
	1.	Setup correct DP table. （状态转移方程）
*/
func coinChangeWithDP(coins []int, amount int) int {

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = MinInt(dp[i-coin]+1, dp[i])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

/**
KP.
	1. Down to top with DP table approach
*/
func coinChangePV2(coins []int, amount int) int {

	// KP.	The meaning of DP.
	// dp[i] the min change ways for amount i.
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1 // KP. Why amount+1?, why not math.MaxInt?
		//dp[i] = math.MaxInt

	}
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 { // amount i < the value of this coin, skip
				continue
			}
			dp[i] = common.MinInt(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
