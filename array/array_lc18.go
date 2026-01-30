package array

import "sort"

func ThreeSum(nums []int, start, target int) [][]int {
	result := [][]int{}

	for idx := start; idx < len(nums)-2; {
		val := nums[idx]
		tuples := TwoSum(nums, idx+1, target-val)
		for _, tuple := range tuples {
			result = append(result, append([]int{val}, tuple...))
		}
		for idx < len(nums)-2 && val == nums[idx] {
			idx++
		}
	}

	return result

}

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	result := [][]int{}

	for idx := 0; idx < len(nums)-3; {
		val := nums[idx]
		tuples := ThreeSum(nums, idx+1, target-val)
		for _, tuple := range tuples {
			result = append(result, append([]int{val}, tuple...))
		}
		for idx < len(nums)-3 && val == nums[idx] {
			idx++
		}
	}

	return result
}
