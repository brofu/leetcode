package dfs

/*
DFS

 1. Time Complexity => WRONG VERSION
    a. For loop to find the starting position O(m*n)
    b. For loop to check if all the empty cells hit O(m*n)
    c. Loop the graph, O(V+E), V=m*n, E = 4*m*n (each node has 4 edges to the other nodes)
    c. Totally O(7*m*n) = O(m*n)

 2. Time Complexity => Correct version
    This is NOT a stand DFS traverse problem, becase, we need explore ALL the possible pathes
    This is a DFS with backtrack
    If there is E walkable cells in grid, then, the time complexity is around O(E!), since, we need
    a. hit ALL the walkable cells and
    b. each of them can only be walked for exactly ONE time.
    Hence, this is similar as a permutation problem

 3. Space Complexity
    The recursive depth would be O(m*n)
*/
func uniquePathsIII(grid [][]int) int {

	result := 0
	startI := 0
	startJ := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				startI = i
				startJ = j
				break
			}
		}
	}
	var dfs func(int, int)

	dfs = func(i, j int) {

		// base case 1
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == -1 {
			return
		}

		// base case 2
		if grid[i][j] == 2 { // arrive the ending
			if checkAllHit(grid) {
				result++
			}
			return
		}

		// choose
		grid[i][j] = -1

		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}

		// cancel
		grid[i][j] = 0
	}

	dfs(startI, startJ)

	return result
}

func checkAllHit(grid [][]int) bool {
	allHit := true
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				allHit = false
				break
			}
		}
	}
	return allHit
}
