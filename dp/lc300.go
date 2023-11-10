package dp

// nums has at least 1 element
func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = MaxInt(dp[i], dp[j]+1)
			}
		}
	}

	res := 1
	for _, re := range dp {
		if res < re {
			res = re
		}
	}

	return res
}

func lengthOfLISWithBS(nums []int) int {

	top := make([]int, len(nums))

	piles := 0

	for _, number := range nums {

		left, right := 0, piles

		for left < right {
			mid := (left + right) / 2
			if top[mid] < number {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == piles {
			piles += 1
		}

		top[left] = number
	}

	return piles
}
