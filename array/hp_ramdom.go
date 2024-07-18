package array

import "sort"

func triplev1(nums []int, target int) int {

	sort.Ints(nums)

	count := 0
	for i := 0; i < len(nums); i++ {
		count += twoSumHR(nums, i+1, target-nums[i])
	}

	return count
}

// wrong version
func twoSumHR(nums []int, start, target int) int {

	low, high := start, len(nums)-1
	count := 0

	for low < high {
		sum := nums[low] + nums[high]
		if sum <= target {
			low++
		} else {
			count += (low - start + 1)
			high--
		}
	}
	return count
}

func triplev2(nums []int, target int) int {

	sort.Ints(nums)

	count := 0
	for i := 0; i < len(nums); i++ {
		count += twoSumHRV2(nums, i+1, target-nums[i])
	}

	return count
}

func twoSumHRV2(nums []int, start, target int) int {

	count := 0
	j := len(nums) - 1

	for i := start; i < j; i++ {
		for ; j > i; j-- {
			if nums[i]+nums[j] <= target {
				count += (j - i)
				break
			}
		}
	}
	return count
}
