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

func QuickSortInt(nums []int) {

	var quickSortHelper func(int, int)
	quickSortHelper = func(low, high int) {
		if low <= high {
			pivot := high
			cursor := low
			for i := low; i <= high; i++ {
				if nums[i] < nums[pivot] {
					nums[i], nums[cursor] = nums[cursor], nums[i]
					cursor++
				}
			}
			pivot = cursor
			nums[cursor], nums[high] = nums[high], nums[cursor]

			quickSortHelper(low, pivot-1)
			quickSortHelper(pivot+1, high)
		}
	}

	quickSortHelper(0, len(nums)-1)
}

func MergeSortPV2(nums []int) {

	tmp := make([]int, len(nums))

	var mergeSortHelper func(int, int)
	mergeSortHelper = func(low, high int) {

		// base case
		if low == high {
			return
		}

		mid := (low + high) / 2
		mergeSortHelper(low, mid)
		mergeSortHelper(mid+1, high)

		// KP. Note the location of this statement
		for i := low; i <= high; i++ {
			tmp[i] = nums[i]
		}

		// merge. postorder logic
		i, j, k := low, mid+1, low

		for i <= mid && j <= high {
			if tmp[i] <= tmp[j] {
				nums[k] = tmp[i]
				i++
			} else {
				nums[k] = tmp[j]
				j++
			}
			k++
		}

		for ; i <= mid; i, k = i+1, k+1 {
			nums[k] = tmp[i]
		}
		for ; j <= mid; j, k = j+1, k+1 {
			nums[k] = tmp[j]
		}

	}
	mergeSortHelper(0, len(nums)-1)
}
