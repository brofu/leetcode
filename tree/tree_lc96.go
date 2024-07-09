package tree

func numTrees(n int) int {

	var count func(int, int) int

	memo := make([][]int, n+1)
	for index := range memo {
		memo[index] = make([]int, n+1)
	}

	count = func(low, high int) int {

		if low >= high {
			return 1
		}

		res := 0

		if res = memo[low][high]; res != 0 {
			return res
		}

		for i := low; i <= high; i++ {
			left := count(low, i-1)
			right := count(i+1, high)
			res += left * right
		}

		memo[low][high] = res

		return res
	}

	return count(1, n)
}

/**
KP.
	Why no need of `[][]int` for memo? ==> if the range length (high-low) are same, then the result are same.
*/
func numTreesPV1(n int) int {

	memo := make([]int, n)
	var count func(int, int) int

	count = func(low, high int) int {
		if low >= high {
			return 1
		}

		// memo
		if memo[high-low] != 0 {
			return memo[high-low]
		}

		result := 0
		for i := low; i <= high; i++ {
			left := count(low, i-1)
			right := count(i+1, high)
			result += left * right
		}
		memo[high-low] = result
		return result
	}

	return count(1, n)
}
