package array

/*
KP:
	1. Pre-sum versions
		max-of-the-presum - min-of-the-presum.
	2. Other versions:
		DP

TC:
	1. One loop. O(n)

SC:
	1. O(1)
*/
func maxSubArray(nums []int) int {

	sum := 0
	minSum := 0
	result := nums[0]

	for _, n := range nums {
		sum += n
		if sum-minSum > result { // the max value would be collected
			result = sum - minSum
		}
		if sum < minSum { // search the min of presum
			minSum = sum
		}
	}

	return result
}
