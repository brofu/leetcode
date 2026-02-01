package array

func productExceptSelf(nums []int) []int {

	n := len(nums)
	lm := make([]int, n)
	rm := make([]int, n)
	ans := make([]int, n)

	lm[0] = nums[0]
	rm[n-1] = nums[n-1]
	for i, j := 1, n-2; i < n && j >= 0; i, j = i+1, j-1 {
		lm[i] = lm[i-1] * nums[i]
		rm[j] = rm[j+1] * nums[j]
	}
	ans[0] = rm[1]
	ans[n-1] = lm[n-2]
	for i := 1; i < n-1; i++ {
		ans[i] = lm[i-1] * rm[i+1]
	}
	return ans
}

/*
KP:
	1. Reuse the `result` array
*/
func productExceptSelfV2(nums []int) []int {

	n := len(nums)
	ans := make([]int, n)

	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	rm := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= rm
		rm *= nums[i]
	}
	return ans
}
