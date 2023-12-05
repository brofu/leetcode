package tree

import (
	"github.com/brofu/leetcode/common"
)

func reversePairs(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	// O(2N)
	sorted := make([]int, len(nums)*2)

	for index, val := range nums {
		sorted[index*2] = val
		sorted[index*2+1] = val * 2
	}

	common.MergeSort(sorted)

	// set up index map
	index := 0
	indexMap := make(map[int]int)
	for _, val := range sorted {
		if _, exists := indexMap[val]; !exists {
			indexMap[val] = index
			index += 1
		}
	}

	bit := NewBinaryIndexTreeReverse(len(indexMap))

	sum := 0
	for index := 0; index < len(nums); index++ {
		val := nums[index]
		sum += bit.Sum(indexMap[val*2] + 1)
		bit.Add(indexMap[val], 1)
	}

	return sum
}

func reversePairsV2(nums []int) int {
	return reversePairsWrapper(nums, 0, len(nums)-1)
}

func reversePairsWrapper(nums []int, left, end int) int {

	if left >= end {
		return 0
	}

	middle := (left + end) / 2
	leftCount := reversePairsWrapper(nums, left, middle)
	rightCount := reversePairsWrapper(nums, middle+1, end)

	count, i, j := 0, left, middle+1

	// count
	for i <= middle && j <= end {
		if nums[i] > 2*nums[j] {
			count += middle - i + 1
			j += 1
		} else {
			i += 1
		}
	}
	// sort
	common.MergeSort(nums[left : end+1])
	return count + leftCount + rightCount
}
