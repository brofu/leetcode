package array

/*
KP.
	1. Get key point. count[sum[i]-k],出现过几次，就有几次以 `i-1`, 结尾的和为k的子数组
*/
func subarraySum(nums []int, k int) int {

	result := 0
	m := make(map[int]int)

	sum := 0
	m[0] = 1
	for i := 1; i < len(nums)+1; i++ {
		sum += nums[i-1]
		need := sum - k
		if c, exists := m[need]; exists {
			result += c
		}
		m[sum]++
	}

	return result
}
