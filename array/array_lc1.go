package array

import "github.com/brofu/leetcode/common"

/**
KP.
	This version only work when there is NO duplicate elements in the array.
*/
func twoSum(nums []int, target int) []int {
	result := []int{}

	records := make(map[int]int, len(nums))
	for index, value := range nums {
		records[value] = index
	}

	common.QuickSort(nums)

	for low, high := 0, len(nums)-1; low < high; {
		sum := nums[low] + nums[high]
		if sum < target {
			low++
		}
		if sum == target {
			lowIndex := records[nums[low]]
			highIndex := records[nums[high]]
			result = append(result, []int{lowIndex, highIndex}...)
			low++
		}
		if sum > target {
			high--
		}
	}
	return result
}

func twoSumRightVersion(nums []int, target int) []int {
	result := []int{}
	records := make(map[int]int)

	for i, v := range nums {
		left := target - v
		if j, ok := records[left]; !ok {
			records[v] = i
		} else {
			result = []int{j, i}
		}
	}
	return result
}
