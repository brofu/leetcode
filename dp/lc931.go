package dp

func minFallingPathSum(matrix [][]int) int {

	maxSumTop := 10001
	n := len(matrix)
	memo := make([][]int, n)

	for i := 0; i < n; i++ {
		memo[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			memo[i][j] = maxSumTop
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		maxSumTop = MinInt(maxSumTop, dpMinFallingPathSum(matrix, memo, n-1, j))
	}

	return maxSumTop
}

func dpMinFallingPathSum(matrix, memo [][]int, i, j int) int {

	maxSumTop := 10001

	// edged case
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return maxSumTop
	}

	// base case
	if i == 0 {
		return matrix[0][j]
	}

	// get from memo
	if memo[i][j] != maxSumTop {
		return memo[i][j]
	}

	// calculate
	memo[i][j] = matrix[i][j] + MinInt(MinInt(dpMinFallingPathSum(matrix, memo, i-1, j-1), dpMinFallingPathSum(matrix, memo, i-1, j)), dpMinFallingPathSum(matrix, memo, i-1, j+1))

	return memo[i][j]
}
