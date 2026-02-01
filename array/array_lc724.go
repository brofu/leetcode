package array

func pivotIndex(nums []int) int {

	n := len(nums)
	sum := make([]int, n+1)
	for idx := 1; idx < len(sum); idx++ {
		sum[idx] = sum[idx-1] + nums[idx-1]
	}
	for idx := 0; idx < n; idx++ {
		if sum[idx] == sum[n]-sum[idx+1] {
			return idx
		}
	}
	return -1
}

/*
KP:
	1. Easier to understand
*/
func pivotIndexV2(nums []int) int {

	n := len(nums)
	sum := make([]int, n+1)
	for idx := 1; idx < len(sum); idx++ {
		sum[idx] = sum[idx-1] + nums[idx-1]
	}
	for idx := 1; idx < len(sum); idx++ {
		if sum[idx-1]-sum[0] == sum[len(sum)-1]-sum[idx] {
			return idx - 1
		}
	}
	return -1
}
