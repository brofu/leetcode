package array

/**
KP
	Utilize Merge Sort
*/

func countSmaller(nums []int) []int {

	count := make([]int, len(nums))

	type pair struct {
		val   int
		index int
	}

	temp := make([]pair, len(nums))
	array := make([]pair, len(nums))

	for i := 0; i < len(nums); i++ {
		array[i] = pair{nums[i], i}
	}

	var mergeSort func([]pair, int, int)

	mergeSort = func(array []pair, low, high int) {
		if low == high {
			return
		}

		mid := (low + high) / 2

		mergeSort(array, low, mid)
		mergeSort(array, mid+1, high)

		for i := low; i <= high; i++ {
			temp[i] = array[i]
		}

		index, i, j := low, low, mid+1
		for ; i <= mid && j <= high; index += 1 {
			if temp[i].val <= temp[j].val { // KP. "<=" instead of "=".
				array[index] = temp[i]
				count[temp[i].index] += j - mid - 1 // KP. How does this work?
				i += 1
			} else {
				array[index] = temp[j]
				j += 1
			}
		}

		for ; i <= mid; index, i = index+1, i+1 {
			array[index] = temp[i]
			count[temp[i].index] += j - mid - 1
		}

		for ; j <= high; index, j = index+1, j+1 {
			array[index] = temp[j]
		}
	}

	mergeSort(array, 0, len(nums)-1)
	return count
}
