package backtrack

import "sort"

/**
KP
	1.	How to pruning? The duplicated elements may lead to duplicate subsets
		*	sort the numbers
		*	skip the number if nums[i] == nums[i-1]
*/
func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	track := []int{}

	sort.Ints(nums)
	var bt func([]int, []int, int, *[][]int)
	bt = func(nums []int, track []int, start int, result *[][]int) {

		temp := make([]int, len(track))
		copy(temp, track)
		*result = append(*result, temp)

		for i := start; i < len(nums); i++ {

			// filter
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// choose, next layer traverse, and cancel choose
			bt(nums, append(track, nums[i]), i+1, result)
		}
	}

	bt(nums, track, 0, &result)
	return result
}
