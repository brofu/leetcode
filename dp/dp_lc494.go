package dp

import (
	"github.com/brofu/leetcode/common"
)

/**
KP.
	1.	How to convert the problem to Knapsack Problem?
		a.	sum(A) - sum(B) = target
		b.	sum(A) * 2 = target + sum(B) + sum(A) = target + sum(nums)
		c.	sum(A) = (target + sum(nums)) / 2
		c.	0-1 Knapsack Problem: How to pick items to fullfil the bag with capability (target + sum(nums)) / 2?
	2.	Base case.
		a.	dp[0][0] = 1. bag capability is 0, and there is NO items to put in
		b.  What's the value of dp[i][0]  when i != 0? Need to calculate, because the weight of an item is possible 0 in this case.
	3.	With space compression.
*/

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < common.AbsIntSub(target, 0) || (sum+target)%2 != 0 {
		return 0
	}

	weight := (sum + target) / 2
	dp := make([]int, weight+1)
	dp[0] = 1 //KP. Why not change this inner the loop?

	// dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
	for i := 1; i <= len(nums); i++ {
		for j := weight; j >= 0; j-- {
			if j >= nums[i-1] {
				// dp[j] is actually dp[i-1][j] before update
				// dp[j-nums[i-1]] is actually dp[i-1][j-nums[i-1]], and that's why need to loop j from big to small.
				dp[j] = dp[j] + dp[j-nums[i-1]]
			} // else,  dp[j] = dp[j]
		}
	}

	return dp[weight]
}
