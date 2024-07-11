package common

/**
KP.
	Reduce the creation of slice objects
*/
func MergeSortV2(nums []int) {

	temp := make([]int, len(nums))

	var mergeSort func([]int, int, int)
	mergeSort = func(nums []int, low, high int) {
		if low == high {
			return
		}
		mid := (low + high) / 2

		// merge left part
		mergeSort(nums, low, mid)
		// merge right part
		mergeSort(nums, mid+1, high)

		// merge the whole part
		for index := low; index <= high; index++ {
			temp[index] = nums[index]
		}

		index, i, j := low, low, mid+1
		for ; i <= mid && j <= high; index += 1 {
			if temp[i] <= temp[j] {
				nums[index] = temp[i]
				i += 1
			} else {
				nums[index] = temp[j]
				j += 1
			}
		}
		for ; i <= mid; index, i = index+1, i+1 {
			nums[index] = temp[i]
		}
		for ; j <= high; index, j = index+1, j+1 {
			nums[index] = temp[j]
		}

	}

	mergeSort(nums, 0, len(nums)-1)
}
