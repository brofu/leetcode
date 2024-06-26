package array

import (
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
	result := 0

	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	calculate := func(start int, target int) int {
		count := 0
		low, high := start, len(nums)-1
		for low < high {
			if nums[low]+nums[high] < target {
				count += high - low
				low++
			} else {
				high--
			}
		}
		return count
	}

	for i := 0; i < len(nums)-2; i++ {
		t := calculate(i+1, target-nums[i])
		result += t
	}
	return result
}
