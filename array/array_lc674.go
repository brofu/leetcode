package array

func findLengthOfLCIS(nums []int) int {

	current, result := 1, 1

	for index := 0; index < len(nums)-1; index++ {
		if nums[index+1] > nums[index] {
			current += 1
		} else {
			if result < current {
				result = current
			}
			current = 1
		}
	}

	if result < current {
		result = current
	}

	return result
}
