package others

/**
This is not DP. Just a for loop
*/
func resolve0730(nums []int) []int {

	record := make(map[int]int)

	for _, n := range nums {
		record[n]++
	}

	dp := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = -1
		for j := i + 1; j < len(nums); j++ {
			fi := record[nums[i]]
			fj := record[nums[j]]
			if fi < fj {
				dp[i] = nums[j]
				break
			}
		}
	}
	return dp
}

func resolve0730Stack(nums []int) []int {

	result := make([]int, len(nums))
	record := make(map[int]int)

	for _, n := range nums {
		record[n]++
	}
	stack := make([]int, len(nums))

	top := 0
	stack[top] = 0

	for index := 1; index < len(nums); index++ {
		if record[nums[stack[top]]] > record[nums[index]] {
			top++
			stack[top] = index
		} else {
			for top >= 0 && record[nums[stack[top]]] < record[nums[index]] {
				result[stack[top]] = nums[index]
				top--
			}
			top++
			stack[top] = index
		}
	}

	for top >= 0 {
		result[stack[top]] = -1
		top--
	}
	return result
}
