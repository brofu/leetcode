package array

func sortColors(nums []int) {

	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	for fast := slow; fast < len(nums); fast++ {
		if nums[fast] == 1 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
