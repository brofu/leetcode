package array

import (
	"github.com/brofu/leetcode/common"
)

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

func twoSumPV1(nums []int, target int) []int {
	result := []int{}
	record := make(map[int]int)

	for index, num := range nums {
		left := target - num
		if v, ok := record[left]; ok {
			result = []int{v, index}
			break
		}
		record[num] = index
	}

	return result
}

func twoSumPV2(nums []int, target int) []int {

	dMap := make(map[int]int)

	for idx, num := range nums {
		left := target - num
		if i, exists := dMap[left]; exists {
			return []int{i, idx}
		} else {
			dMap[num] = idx
		}
	}
	return []int{}
}

/*
KP:
	1. nums should be storted
	2. use double-pointer,

TC:
	1. O(N)

*/
func TwoSum(nums []int, start, target int) [][]int {

	result := [][]int{}

	lo, hi := start, len(nums)-1

	for lo < hi {
		left, right := nums[lo], nums[hi]
		sum := left + right

		if sum < target {
			if lo < hi && left == nums[lo] {
				lo++
			}
			continue
		}
		if sum > target {
			for lo < hi && right == nums[hi] {
				hi--
			}
			continue
		}
		result = append(result, []int{left, right})

		for lo < hi && left == nums[lo] {
			lo++
		}

		for lo < hi && right == nums[hi] {
			hi--
		}
	}
	return result
}
