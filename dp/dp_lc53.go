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
