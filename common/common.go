package common

//Comparable is the interface which is comparable
type Comparable interface {
	// if a > b, a.CompareTo(b) return 1
	// if a == b, a.CompareTo(b) return 0
	// if a < b, a.CompareTo(b) return -1
	CompareTo(Comparable) int
}

// Merge Sort

func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	middle := len(nums) / 2

	MergeSort(nums[:middle])
	MergeSort(nums[middle:])

	left := make([]int, middle)
	right := make([]int, len(nums)-middle)

	_ = copy(left, nums[:middle])
	_ = copy(right, nums[middle:])

	index, i, j := 0, 0, 0
	for ; i < middle && j < len(nums)-middle; index += 1 {
		if left[i] <= right[j] {
			nums[index] = left[i]
			i += 1
		} else {
			nums[index] = right[j]
			j += 1
		}
	}

	for ; i < middle; index, i = index+1, i+1 {
		nums[index] = left[i]
	}
	for ; j < len(nums)-middle; index, j = index+1, j+1 {
		nums[index] = right[j]
	}
}

func MergeSortPV1(nums []int) {
	if len(nums) <= 1 {
		return
	}

	middle := len(nums) / 2

	MergeSortPV1(nums[:middle])
	MergeSortPV1(nums[middle:])

	left := make([]int, len(nums[:middle]))
	right := make([]int, len(nums[middle:]))

	_ = copy(left, nums[:middle])
	_ = copy(right, nums[middle:])

	i, j := 0, 0
	index := 0

	for ; i < middle && j < len(nums[middle:]); index += 1 {
		if left[i] < right[j] {
			nums[index] = left[i]
			i += 1
		} else {
			nums[index] = right[j]
			j += 1
		}
	}

	for ; i < middle; i, index = i+1, index+1 {
		nums[index] = left[i]
	}

	for ; j < len(nums[middle:]); j, index = j+1, index+1 {
		nums[index] = right[j]
	}
}
