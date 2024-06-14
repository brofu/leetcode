package backtrack

import "sort"

/**
KP
	1.	How to do pruning?
*/
func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))

	sort.Ints(nums)

	var bt func([]int, []int, []bool, *[][]int)

	bt = func(nums []int, track []int, used []bool, result *[][]int) {

		// base case
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			*result = append(*result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// pruning
			if used[i] {
				continue
			}

			// pruning
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// choose, next layer and cancel choose
			used[i] = true
			bt(nums, append(track, nums[i]), used, result)
			used[i] = false
		}
	}

	bt(nums, track, used, &result)
	return result
}
