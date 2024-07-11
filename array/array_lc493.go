package array

func reversePairs(nums []int) int {
	result := 0
	temp := make([]int, len(nums))

	var mergeSort func([]int, int, int)

	mergeSort = func(nums []int, low, high int) {

		if low == high {
			return
		}

		mid := (low + high) / 2
		mergeSort(nums, low, mid)
		mergeSort(nums, mid+1, high)

		for i := low; i <= high; i++ {
			temp[i] = nums[i]
		}

		// get the result
		// KP. The logic is inserted into merge sort
		j := mid + 1
		for i := low; i <= mid; i++ {
			for ; j <= high && temp[i] > 2*temp[j]; j += 1 { // KP. For a value of j, if temp[i] > 2*temp[j], then, temp[i+1] > 2*temp[j]
			}
			result += j - mid - 1
		}

		// merge the complete
		i, index := low, low
		for j = mid + 1; i <= mid && j <= high; index += 1 {
			if temp[i] > temp[j] {
				nums[index] = temp[j]
				j += 1
			} else {
				nums[index] = temp[i]
				i += 1
			}
		}
		for ; i <= mid; i, index = i+1, index+1 {
			nums[index] = temp[i]
		}
		for ; j <= high; j, index = j+1, index+1 {
			nums[index] = temp[j]
		}
	}

	mergeSort(nums, 0, len(nums)-1)
	return result
}
