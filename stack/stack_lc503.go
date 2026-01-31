package stack

/*
KP:
	1. Monotonic Stack
	2. Cycle array ==> for loop with 2*n-1, and with %
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	s := make([]int, 0)

	for i := 2*n - 1; i >= 0; i-- {

		for len(s) > 0 && s[len(s)-1] <= nums[i%n] {
			s = s[:len(s)-1] // pop
		}

		if len(s) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = s[len(s)-1] // peak
		}
		s = append(s, nums[i%n]) // push
	}

	return res
}
