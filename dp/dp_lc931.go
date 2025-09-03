package dp

/*
TC:
1. Initial dp[0], O(N)
2. Update dp[1:], O((M-1)*N), around O(M*N)
3. For each cell, we may need to 3 previous neighbors, that's O(3*MN)
4. Get the find result from dp[m-1], O(N)
5. Over all, it's O(M*N)

SC:
1. dp: O(M*N)
5. Over all, it's O(M*N)
*/
func minFallingPathSumDPA(matrix [][]int) int {

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	for idx, val := range matrix[0] {
		dp[0][idx] = val
	}

	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		v1, v2, v3 := 10001, 10001, 10001
		for j := 0; j < n; j++ {
			if j-1 >= 0 {
				v1 = dp[i-1][j-1]
			}
			v2 = dp[i-1][j]
			if j+1 < n {
				v3 = dp[i-1][j+1]
			}
			dp[i][j] = MinIntTrix(v1, v2, v3) + matrix[i][j]
		}
	}
	result := dp[m-1][0]
	for i := 1; i < n; i++ {
		result = MinInt(result, dp[m-1][i])
	}

	return result
}

/*
TC:
1. Construct memo: O(M*N)
2. Recursive call: Depth, O(M-1), at each layer, we need to check 3 neighbors, around O(3*(M-1))
3. Fort All the numbers of the last row, need to trigger the recursive call. that's O(N*3*(M-1))
4. Overall, it's O(M*N) + (3*N*(M-1)) = O(M*N)

SC:
1. Memo: O(MN)
*/
func minFallingPathSumDPF(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -10001
		}
	}

	var dp func(int, int) int

	dp = func(i, j int) int {

		// base case
		if i < 0 || i >= m || j < 0 || j >= n {
			return 10001
		}
		if memo[i][j] != -10001 {
			return memo[i][j]
		}
		// base case
		if i == 0 {
			return matrix[0][j]
		}

		memo[i][j] = MinIntTrix(dp(i-1, j-1), dp(i-1, j), dp(i-1, j+1)) + matrix[i][j]
		return memo[i][j]
	}

	result := 10001
	for idx := range matrix[m-1] {
		result = MinInt(result, dp(m-1, idx))
	}
	return result
}

/*
KP
1. dfs version
2. Actually, it's a 3-nary tree, with height as m
*/
func minFallingPathSumDFS(matrix [][]int) int {

	m, n := len(matrix), len(matrix[0])
	result := 10001

	var dfs func(i, j, sum int)
	dfs = func(i, j, sum int) {

		//base case
		if i == m-1 {
			result = MinInt(result, sum)
		}

		for _, neighbor := range [][]int{{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1}} {
			// pruning
			x, y := neighbor[0], neighbor[1]
			if x >= m || y < 0 || y >= n {
				continue
			}
			dfs(x, y, sum+matrix[x][y])
		}
	}

	for idx, val := range matrix[0] {
		dfs(0, idx, val)
	}

	return result
}
