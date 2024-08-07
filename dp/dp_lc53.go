package dp

import (
	"fmt"
	"math"

	"github.com/brofu/leetcode/common"
)

func maxSubArray(nums []int) int {

	result := math.MinInt
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(dp); i++ {
		dp[i] = common.MaxInt(nums[i], dp[i-1]+nums[i])
	}

	fmt.Println("flag", dp)
	for _, num := range dp {
		if result < num {
			result = num
		}
	}

	return result
}

/**
KP.
	DP table. 以 nums[i] 为结尾的「最大子数组和」为 dp[i]。
	这种定义之下，想得到整个 nums 数组的「最大子数组和」，不能直接返回 dp[n-1]，而需要遍历整个 dp 数组：
*/
func maxSubArrayPV1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]

	// dp[i] is the max sum ends with nums[i]
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = common.MaxInt(nums[i], dp[i-1]+nums[i])
	}

	for _, sum := range dp {
		if result < sum {
			result = sum
		}
	}

	return result
}

/**
前缀和数组 preSum 就是 nums 元素的累加和，preSum[i+1] - preSum[j] 其实就是子数组 nums[j..i] 之和（根据 preSum 数组的实现，索引 0 是占位符，所以 i 有一位索引偏移）。
那么反过来想，以 nums[i] 为结尾的最大子数组之和是多少？其实就是 preSum[i+1] - min(preSum[0..i])。
*/
func maxSubArrayPreSum(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	minSum := math.MaxInt
	result := nums[0]
	for i := 1; i < len(preSum); i++ {
		minSum = common.MinInt(minSum, preSum[i-1])
		result = common.MaxInt(preSum[i]-minSum, result)
	}

	return result
}
