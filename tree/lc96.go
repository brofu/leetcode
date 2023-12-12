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
