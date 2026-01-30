package dp

/*
KP.
	1. This problem can also be resolved by backtrack
	2. This equals to, select numbers from []int, of which the sum is amount. Each number can be used for multiple times
	3. But the TC would be high


TC:
	1. The recursive depth is O(M). M is the length of coins (or the coins number). N around to A/C. A is the amount, and C is minimum coin
	2. At each node, there is a for loop. Worst is O(M)
	3. No other operation (result++, sum+coins[i], all O(1)), so over all TC: O(M*M^(A/C))

SC:
	1. Recursive depth, O(M)
*/

func changeBackTrack(amount int, coins []int) int {

	var (
		result = 0
		bt     func(int, int)
	)

	bt = func(idx, sum int) {

		//base case
		if sum == amount {
			result++
			return
		}

		// base case 2
		if sum > amount {
			return
		}

		for i := idx; i < len(coins); i++ {
			// choose, explore, and cancel
			bt(i, sum+coins[i])
		}
	}

	bt(0, 0)
	return result
}

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

func changeDP(amount int, coins []int) int {

	var (
		n = len(coins)
		// dp[i][j]: answers with items [0....i], and amount as `j`.
		dp = make([][]int, n+1)
	)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] { // cannot put coins[i] into bag
				dp[i][j] = dp[i-1][j]
			} else { // can choose put coins[i] into bag or NOT
				dp[i][j] = dp[i-1][j] + // not put
					dp[i][j-coins[i-1]] // put (note, coin can be used for multiple time)
			}
		}
	}
	return dp[n][amount]
}

/*
KP:
	1. The order of update dp[j]. Why should loop from coins[i-1] to amount?
*/
func changeDPCompressed(amount int, coins []int) int {

	var (
		dp = make([]int, amount+1)
	)

	dp[0] = 1

	for i := 1; i < len(coins)+1; i++ {
		for j := coins[i-1]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i-1]]
		}
	}

	return dp[amount]
}
