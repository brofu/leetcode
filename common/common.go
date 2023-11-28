package common

import "fmt"

//Comparable is the interface which is comparable
type Comparable interface {
	// if a > b, a.CompareTo(b) return 1
	// if a == b, a.CompareTo(b) return 0
	// if a < b, a.CompareTo(b) return -1
	CompareTo(Comparable) int
}

// Merge Sort

func MergeSort(nums []int) {
	mergeSortWrapper(nums, 0, len(nums)-1)
}

func mergeSortWrapper(nums []int, start, end int) {
	size := end - start

	if size <= 1 {
		return
	}

	middle := size / 2

	// divide
	mergeSortWrapper(nums, start, middle)
	mergeSortWrapper(nums, middle+1, end)

	// conquer

	left := nums[start : middle+1]
	right := nums[middle+1 : end+1]

	i, j, k := start, 0, 0

	fmt.Println(nums, start, middle, middle+1, end, left, right)
	for ; j <= middle && k < end-middle; i += 1 {
		fmt.Println(i, j, k)
		if left[j] < right[k] {
			nums[i] = left[j]
			j += 1
		} else {
			nums[i] = right[k]
			k += 1
		}
	}

	for ; j <= middle; i, j = i+1, j+1 {
		nums[i] = left[j]
	}

	for ; k <= end; i, k = i+1, k+1 {
		nums[i] = right[k]
	}

}
