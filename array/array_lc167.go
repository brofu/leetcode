package array

func twoSum167(numbers []int, target int) []int {

	result := []int{}
	for left, right := 0, len(numbers)-1; left <= right; {
		sum := numbers[left] + numbers[right]
		if sum == target {
			result = []int{left + 1, right + 1}
			break
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}

	return result
}
