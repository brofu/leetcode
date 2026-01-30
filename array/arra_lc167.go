package array

func twoSum167PV1(numbers []int, target int) []int {

	lo, hi := 0, len(numbers)-1

	for lo < hi {
		left, right := numbers[lo], numbers[hi]
		sum := left + right

		if sum < target {
			for lo < hi && left == numbers[lo] {
				lo++
			}
			continue
		}

		if sum > target {
			for lo < hi && right == numbers[hi] {
				hi--
			}
			continue
		}
		return []int{lo + 1, hi + 1}
	}

	return nil
}
