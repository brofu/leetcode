package array

import (
	"sort"

	"github.com/brofu/leetcode/common"
)

func isPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i <= j; {
		formatI, okI := common.CheckAndFormatAlphanumeric(s[i])
		if !okI {
			i++
		}
		formatJ, okJ := common.CheckAndFormatAlphanumeric(s[j])
		if !okJ {
			j--
		}

		if okI && okJ {
			if formatI != formatJ {
				return false
			}
			i++
			j--
		}
	}

	return true
}

/*

TC:
	1. sort, O(N*lgN)
	2. out-layer for loop, O(N), for each, call TwoSum, with O(N). So, O(N^2)
	3. Overall O(N^2)
*/
func threeSum(nums []int) [][]int {

	results := [][]int{}
	sort.Ints(nums)

	for idx := 0; idx < len(nums)-2; {
		val := nums[idx]
		tuples := TwoSum(nums, idx+1, 0-val)
		for _, tuple := range tuples {
			results = append(results, append([]int{nums[idx]}, tuple...))
		}
		for idx < len(nums)-2 && val == nums[idx] {
			idx++
		}
	}
	return results
}
