package array

/**
KP.
	1.	Convert the problem to `preSum` of array
		sum of array C(i). C(i) = count(j) where lower <= presum[j] - presum[i] <= upper
	2.	Utilize the `merge sort`.
		*	`Divide and Conquer`.
		*	Keep the relative position in the array. Handle the logic before `merge the whole part`
	3.	Double pointer for ordered array for optimization
*/
func countRangeSum(nums []int, lower int, upper int) int {

	result := 0

	// KP.
	// calculate the presum
	preSum := make([]int, len(nums)+1)
	for index, val := range nums {
		preSum[index+1] = val + preSum[index]
	}

	// the temp space
	temp := make([]int, len(preSum))
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

		// KP. Insert the logic in the merge sort processing. To keep the relative position info.
		// the logic
		start, end := mid+1, mid+1
		for i := low; i <= mid; i++ { // KP. Double Pointer for optimization
			for start <= high && temp[start]-temp[i] < lower {
				start += 1
			}
			for end <= high && temp[end]-temp[i] <= upper {
				end += 1
			}
			result += end - start
		}

		// merge the whole part
		i, j, index := low, mid+1, low
		for ; i <= mid && j <= high; index += 1 {
			if temp[i] < temp[j] {
				nums[index] = temp[i]
				i += 1
			} else {
				nums[index] = temp[j]
				j += 1
			}
		}

		for ; i <= mid; i, index = i+1, index+1 {
			nums[index] = temp[i]
		}
		for ; j <= high; j, index = j+1, index+1 {
			nums[index] = temp[j]
		}
	}

	mergeSort(preSum, 0, len(preSum)-1)
	return result
}
