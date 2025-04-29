package dijkstra

import (
	"container/heap"
	"math"

	"github.com/brofu/leetcode/common"
	myHeap "github.com/brofu/leetcode/common/heap"
)

func minimumEffortPath(heights [][]int) int {

	m, n := len(heights), len(heights[0])

	effortToNode := make([][]int, m)
	for i := 0; i < m; i++ {
		effortToNode[i] = make([]int, n)
		for j := 0; j < n; j++ {
			effortToNode[i][j] = math.MaxInt
		}
	}

	effortToNode[0][0] = 0

	hp := myHeap.NewInterfacePriorityQueue(func(data []interface{}, i, j int) bool {
		return data[i].([]int)[2] < data[j].([]int)[2]
	})

	heap.Push(hp, []int{0, 0, 0})

	for hp.Len() > 0 {
		current := heap.Pop(hp).([]int)

		currentX, currentY, currentDist := current[0], current[1], current[2]

		if currentDist > effortToNode[currentX][currentY] { // exists a smaller distance
			continue
		}

		effortToNode[currentX][currentY] = currentDist

		if currentX == m-1 && currentY == n-1 { // already arrived at the bottom-right node
			continue
		}

		// check neighbors
		for _, neighbor := range [][]int{{currentX - 1, currentY}, {currentX, currentY - 1}, {currentX + 1, currentY}, {currentX, currentY + 1}} {
			neighborX, neighborY := neighbor[0], neighbor[1]
			if neighborX < 0 || neighborX == m || neighborY < 0 || neighborY == n { //invalid neighbors
				continue
			}
			// KEY POINT
			// The definition and update logic of `effortToNode`
			effortToNeighbor := common.MaxInt(effortToNode[currentX][currentY], common.AbsIntSub(heights[currentX][currentY], heights[neighborX][neighborY]))
			if effortToNeighbor < effortToNode[neighborX][neighborY] {
				effortToNode[neighborX][neighborY] = effortToNeighbor
				heap.Push(hp, []int{neighborX, neighborY, effortToNeighbor})
			}
		}
	}

	return effortToNode[m-1][n-1]
}
