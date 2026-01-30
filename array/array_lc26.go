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

/*
TC:	O(n)
*/
func removeDuplicatesPV1(nums []int) int {

	i, j := 0, 1

	for i < len(nums)-1 { // do we really need this?
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		if j == len(nums) {
			return i + 1
		}
		i++
		nums[i] = nums[j]
	}

	return i + 1
}
