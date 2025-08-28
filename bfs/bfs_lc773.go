package bfs

import (
	"bytes"
	"fmt"
	"strings"
)

/*
Key Points

* TC

 1. build board_str, O(M*N)
 2. build neighbor index map, O(M*N*4)
 3. BFS around O(T!), T = M*N
 4. at each BFS, need to
    a. compare if current == target O(M*N)
    b. get the index of 0, that's O(M*N)
 5. So, over all is O((M*N)!*(M*N))

* SC

1. visited map, O(T!), T = M*N? Is this right? NO. There is T! keys, and each key, the length is T. So, this part is O(T!*T)
2. neighbor index map, around O(T*4)
3. queue, around O(T!)
5. So, overall is O(2*(M*N)!*(M*N)+4M*N)
*/
func slidingPuzzle(board [][]int) int {

	visited := make(map[string]struct{})

	var (
		b        bytes.Buffer
		boardStr string
		result   int
		target   string = "123450"
		m        int    = len(board)
		n        int    = len(board[0])
	)

	for _, rows := range board {
		for _, column := range rows {
			b.Write([]byte(fmt.Sprintf("%d", column)))
		}
	}
	boardStr = b.String()

	// construct the neighbor index map
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	neighberIndexMap := make([][]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			neighbers := []int{}
			for _, direction := range directions {
				x := i + direction[0]
				y := j + direction[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				neighbers = append(neighbers, x*n+y)
			}
			neighberIndexMap[i*n+j] = neighbers
		}
	}

	q := []string{boardStr}

	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {

			current := q[idx]

			if current == target {
				return result
			}
			neighbers := getNeighbers(neighberIndexMap, q[idx])

			for _, neighber := range neighbers {
				if _, exists := visited[neighber]; exists {
					continue
				}
				q = append(q, neighber)
				visited[neighber] = struct{}{}
			}
		}
		result++
		q = q[l:]
	}
	return -1
}

func getNeighbers(neighberMap [][]int, current string) []string {

	idx := strings.Index(current, "0")

	neighberIndexes := neighberMap[idx]

	neighbers := make([]string, len(neighberIndexes))
	for i, neighberIndex := range neighberIndexes {
		neighbers[i] = swap(current, idx, neighberIndex)
	}

	return neighbers
}

func swap(str string, s, d int) string {
	l := strings.Split(str, "")
	l[s], l[d] = l[d], l[s]
	return strings.Join(l, "")
}
