package array

func checkSubarraySum(nums []int, k int) bool {

	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	m := make(map[int]int)
	for i := 0; i < n+1; i++ {
		if idx, exists := m[preSum[i]%k]; !exists {
			m[preSum[i]%k] = i
		} else {
			if i-idx >= 2 {
				return true
			}
		}
	}
	return false
}

/*
KP
	1. A better solution, less space complexity
*/
func checkSubarraySumV2(nums []int, k int) bool {

	n := len(nums)
	m := make(map[int]int)
	sum := 0
	m[0] = 0
	for i := 1; i < n+1; i++ {
		sum += nums[i-1]
		r := sum % k
		if idx, exists := m[r]; !exists {
			m[r] = i
		} else {
			if i-idx > 1 {
				return true
			}
		}
	}
	return false
}
