package dp

/**
KP.
	1.	Difference from 0-1 Knapsack Problem?
		a.	No number limitation of each coin
		b.	The status transfer is Different from the 0-1 problem.
*/

func change(amount int, coins []int) int {

	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1 // base case 1. if amount = 0, 1 solution is to do nothing.
	}
	for j := 0; j <= amount; j++ {
		dp[0][j] = 0 // base case 1. if no coin at all, no solution to get the amount
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 { // bag space NOT enough
				dp[i][j] = dp[i-1][j]
			} else { // pick up coins[i-1] AND NOT pick up coins[i-1]. All possible solutions
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}

/**
Space Compression Version.
KP.
	1.	When to update dp[0]?
*/
func changeWithSpaceCompression(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 0 // KP. dp[i][0] = 0, if i = 0
	for i := 1; i <= len(coins); i++ {
		dp[0] = 1 // KP. dp[i][0] = 1, if i > 0.
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 { // KP. if j-coins[i-1] < 0, dp[j] = dp[j]
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}
	return dp[amount]
}
