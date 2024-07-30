package array

/**
KP
	1.	How to control the most right or most right index?
	2.	Check the edged case. Target doesn't exist in the array.
*/
func searchRange(nums []int, target int) []int {
	left := leftMost(nums, target)
	right := rightMost(nums, target)
	return []int{left, right}
}

func leftMost(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target { // shrunk the range from right side
			right = mid - 1
			continue
		}
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}

	if left >= len(nums) { // it's not possible that left < 0
		return -1
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func rightMost(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target { // shrunk the range from left side
			left = mid + 1
			continue
		}
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}

	if right < 0 { // it's not possible that, right >= len(nums)
		return -1
	}

	if nums[right] == target {
		return right
	}

	return -1
}
