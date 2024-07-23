package array

import (
	"github.com/brofu/leetcode/common"
)

/**
KP.
	*	Hard point is, to convert the problem to the Sliding Window scenarios:
		Get the min sum of a sub-array with length of n - k. n = len(s)
	*	Understand the original problem

*/
func maxScore(cardPoints []int, k int) int {

	sum, n := 0, len(cardPoints)

	for _, v := range cardPoints {
		sum += v
	}

	if n == k {
		return sum
	}

	subLength := n - k
	subMin := sum
	subSum := 0

	for left, right := 0, 0; right < n; right++ {
		subSum += cardPoints[right]
		if right-left+1 == subLength {
			subMin = common.MinInt(subMin, subSum)
			subSum -= cardPoints[left]
			left++
		}
	}
	return sum - subMin
}
