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

func QuickSort(nums []int) {

	var quickSortHelper func([]int, int, int)

	quickSortHelper = func(nums []int, low, high int) {
		if low < high {

			// similar to pre-order of bt
			p := high
			v := nums[p]
			cursor := low
			for i := low; i < high; i++ {
				if nums[i] < v {
					nums[i], nums[cursor] = nums[cursor], nums[i]
					cursor++
				}
			}
			nums[cursor], nums[high] = nums[high], nums[cursor]
			p = cursor

			// left tree
			quickSortHelper(nums, low, p-1)
			// right tree
			quickSortHelper(nums, p+1, high)
		}
	}

	quickSortHelper(nums, 0, len(nums)-1)
}
