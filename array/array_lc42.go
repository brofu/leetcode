package array

import "github.com/brofu/leetcode/common"

func trap(height []int) int {

	result := 0
	lMax, rMax := 0, 0
	left, right := 0, len(height)-1

	for left < right {
		lMax = common.MaxInt(lMax, height[left])
		rMax = common.MaxInt(rMax, height[right])

		if lMax < rMax {
			result += lMax - height[left]
			left++
		} else {
			result += rMax - height[right]
			right--
		}
	}

	return result
}
