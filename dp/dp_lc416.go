package dp

import "github.com/brofu/leetcode/common"

/*

KP:
	1. DP function definition: dp(n, w)
	2. Very similar to backtrack

	3. Important!  Wrong! It's similar as a permutation problem.

*/
func BagProblemDPFunction(n, w int, val, weight []int) int {

	var (
		dp   func(int, int) int
		used = make([]bool, n)
	)

	dp = func(n, w int) int {

		if n == 0 || w == 0 {
			return 0
		}

		if w < 0 {
			return -1
		}

		temp := 0
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			r := dp(n-1, w-weight[i])
			if r == -1 {
				continue
			}
			if temp < r+val[i] {
				temp = r + val[i]
			}
			used[i] = false
		}

		return temp
	}

	return dp(n, w)
}

/*

KP:
	1. Back Track version to resolve the problem.
	2. Use the order

TC:
	1. With memo, memo[][], the TC is around O(n^2*w)
	2. What's the difference between this with the DP version?
		本质区别就是：你把“决策点”放在哪。
		DP（选/不选）：每个状态只对“当前这一件物品”做常数个决策
		→ 每个状态只有 2 条转移边（take / skip）
		→ 每状态工作量 O(1)

		你这个 BT-for-loop（枚举下一件）：每个状态要在“剩余所有物品”里挑下一件
		→ 每个状态有 (n−idx) 条转移边
		→ 为了找到最优，你必须把这些候选都看一遍
		→ 每状态工作量最坏 O(n)

		这就是那个额外O(n) 因子从哪里来的：每个状态的出边数量不同。
*/
func BagProblemBackTrack(n, w int, val, weight []int) int {

	var (
		bt     func(int, int, int)
		result int
	)

	bt = func(idx, sumV int, sumW int) {

		// base case 1
		if sumW > w {
			return
		}

		if result < sumV {
			result = sumV
		}

		for i := idx; i < n; i++ {
			// choose, explore, cancel-choose
			bt(i+1, sumV+val[i], sumW+weight[i])
		}
	}

	bt(0, 0, 0)
	return result
}

/*

KP:
	1. The standard DP table version of the 0-1 bag problem
	2. dp[i][j] means, the max value of the (1...i) items and weight j.
	3. The DP table definition

TC:
	1. Overall O(n*w)

SC:
	1. Overall O(n*w), for the dp table
*/
func BagProblemDPTable(n, w int, val, weight []int) int {

	var (
		dp = make([][]int, n+1)
	)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-weight[i-1] < 0 { // cannot put item i, it it's weigh > j
				dp[i][j] = dp[i-1][j]
				continue
			}
			dp[i][j] = common.MaxInt(
				dp[i-1][j], // choose not to put item
				dp[i-1][j-weight[i-1]]+val[i-1],
			)
		}
	}

	return dp[n][w]
}

/*

KP:
	1. Why `j` should looped backward?

SC:
	1. Over all O(w)
*/
func BagProblemDPTableCompressed(n, w int, val, weight []int) int {

	var (
		dp = make([]int, w+1)
	)
	for i := 1; i <= n; i++ {
		for j := w; j >= weight[i-1]; j++ {
			dp[j] = common.MaxInt(
				dp[j], // choose not to put item
				dp[j-weight[i-1]]+val[i-1],
			)
		}
	}
	return dp[w]
}

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

func canPartitionDPV1(nums []int) bool {

	var (
		sum int
	)

	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	n := sum / 2
	// if the is answer with the `first i` items, and `j weights`
	dp := make([][]bool, len(nums)+1)

	for idx := range dp {
		dp[idx] = make([]bool, n+1)
	}
	for idx := range dp { // dp[0...i][0] = true
		dp[idx][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for w := 1; w <= n; w++ {
			if nums[i-1] > w { // cannot put into bag
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = dp[i-1][w] || // choose not to put into bag
					dp[i-1][w-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][n]
}

/*

TC:
	1. O(m*n). m = len(nums), n = sum(nums) / 2

SC:
	1. O(n), n = sum(nums) / 2 + 1
*/
func canPartitionDPCompressedV1(nums []int) bool {
	var (
		sum int
	)

	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	var (
		n  = sum / 2
		dp = make([]bool, n+1)
	)
	dp[0] = true

	for i := 1; i < len(nums)+1; i++ {
		for w := n; w >= nums[i-1]; w-- {
			dp[w] = dp[w] || dp[w-nums[i-1]]
		}
	}

	return dp[n]
}
