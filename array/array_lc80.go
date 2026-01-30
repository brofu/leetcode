package array

/*
SC:
	O(n). Better solution? Noted the array is sorted
*/
func removeDuplicates80(nums []int) int {

	m := make(map[int]int)
	slow, fast := 0, 0

	for fast < len(nums) {
		num := nums[fast]
		if m[num] < 2 {
			m[num]++
			nums[slow] = num
			slow++
		}
		fast++
	}

	return slow
}

/*
KP:
	1. Utilize the sorted array. If nums[fast] == nums[slow-2], it must be the 3+ time to happen
*/
func removeDuplicates80V2(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}

	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
