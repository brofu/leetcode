package array

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	low, fast := 0, 0
	for fast < len(nums) {
		if nums[low] == nums[fast] {
			fast++
			continue
		}
		low++
		nums[low] = nums[fast]
	}

	return low + 1
}
