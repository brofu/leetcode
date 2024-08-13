package dp

/**
KP.
	1.	How to convert the problem to `Knapsack Problem`?
		给一个可装载重量为 sum / 2 的背包和 N 个物品，每个物品的重量为 nums[i]。现在让你装物品，是否存在一种装法，能够恰好将背包装满？
*/
func canPartition(nums []int) bool {

	sum := 0
	for _, w := range nums {
		sum += w
	}

	if sum%2 != 0 { // cannot part
		return false
	}

	w := sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, w+1)
		dp[i][0] = true // base case 1
	}

	for j := 0; j <= w; j++ { // base case 2
		dp[0][j] = false
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= w; j++ {
			if j-nums[i-1] < 0 { // bag space not enough. Not put
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]] // Not put nums[i] into bag or Yes
			}
		}
	}
	return dp[len(nums)][w]
}

/**
KP.
	1.	Observe the 2-D DP table version, dp[i][j] is only depending on dp[i-1][j] and dp[i-1][j-nums[i]]
	2.	How to reuse the result of dp[i-1][...] to use 1-D table?
		a.	Before dp[j] is updated, it's actually dp[i-1][j].
		b.	How to use dp[i-1][j-nums[i]] before it's updated? Loop j from high to low
*/

func canPartitionWithSpaceCompression(nums []int) bool {
	sum := 0
	for _, w := range nums {
		sum += w
	}

	if sum%2 != 0 { // cannot part
		return false
	}

	w := sum / 2

	dp := make([]bool, w+1)
	dp[0] = true

	for i := 1; i <= len(nums); i++ {
		for j := w; j > 0; j-- { // KP. loop from big to small.
			//if j-nums[i-1] < 0 {
			//	dp[j] = dp[j]
			//} else {
			if j-nums[i-1] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i-1]] // dp[j] is actually dp[i-1][p] before update
			}
			//}
		}
	}

	return dp[w]
}
