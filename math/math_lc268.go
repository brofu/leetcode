package math

/*
KP:
	1. Idea of ^
	2. sum(0...n) - sum(nums)
*/
func missingNumber(nums []int) int {

	res := 0
	res ^= len(nums)
	for idx, x := range nums {
		res ^= idx
		res ^= x
	}

	return res
}
