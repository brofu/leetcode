package array

func moveZeroes(nums []int) {

	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroesPV1(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ { // similar as the logic in pivot func in quick sort
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
